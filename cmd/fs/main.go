package main

import (
	"context"
	"estore/api/middlewares"
	"net/http"

	"estore/configs"
	"estore/protoc/fspb"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	api "estore/api/v1"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	configs.Develop()
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigs
		cancel()
	}()

	go httpServer(
		grpcServer(),
	)

	<-ctx.Done()
}

func httpServer(gs *grpc.Server) {
	router := gin.Default()
	router.Use(
		middlewares.CORS(),
		middlewares.GrpcWebProxy(gs),
	)
	api.Http(router)

	httpPort := "8080"
	if os.Getenv("HTTP_PORT") != "" {
		httpPort = os.Getenv("HTTP_PORT")
	}

	router.NoRoute(func(ctx *gin.Context) {
		http.ServeFile(ctx.Writer, ctx.Request, "noroute.html")
	})

	fmt.Println("http service listen at: " + httpPort)
	router.Run(":" + httpPort)
}

func grpcServer() *grpc.Server {
	// creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	// if err != nil {
	// 	panic(err)
	// }

	// gs := grpc.NewServer(grpc.Creds(creds))
	gs := grpc.NewServer()
	fspb.RegisterFileSystemServer(gs, &api.FSService{})
	reflection.Register(gs)

	grpcPort := "8090"
	if os.Getenv("GRPC_PORT") != "" {
		grpcPort = os.Getenv("GRPC_PORT")
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		panic(err)
	}

	fmt.Println("resource service listen at : " + grpcPort)
	go func() {
		if err := gs.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return gs
}
