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

	"go.elastic.co/apm/module/apmgin/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	os.Setenv("ELASTIC_APM_SERVICE_NAME", "estore")
	os.Setenv("ELASTIC_APM_API_KEY", "")
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://localhost:8200")
	os.Setenv("ELASTIC_APM_LOG_LEVEL", "debug")
	os.Setenv("ELASTIC_APM_LOG_FILE", "stderr")

	os.Setenv("ELASTIC_APM_TRANSACTION_SAMPLE_RATE", "1.0")
	os.Setenv("ELASTIC_APM_CAPTURE_HEADERS", "true")
	os.Setenv("ELASTIC_APM_CAPTURE_BODY", "all")
	os.Setenv("ELASTIC_APM_EXIT_SPAN_MIN_DURATION", "1ms")
	os.Setenv("ELASTIC_APM_SPAN_STACK_TRACE_MIN_DURATION", "1ms")
	os.Setenv("ELASTIC_APM_METRICS_INTERVAL", "3s")

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
		apmgin.Middleware(router),
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
	gs := grpc.NewServer(
	// grpc.ChainUnaryInterceptor(apmgrpc.NewUnaryServerInterceptor()),
	// grpc.ChainStreamInterceptor(apmgrpc.NewStreamServerInterceptor()),
	)
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
