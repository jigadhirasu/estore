package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func GrpcWebProxy(gs *grpc.Server) gin.HandlerFunc {

	wrappedGrpc := grpcweb.WrapServer(gs,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	// grpcwebproxy
	return func(ctx *gin.Context) {
		if wrappedGrpc.IsGrpcWebRequest(ctx.Request) {
			// 只要是grpc-web请求，就返回200
			ctx.Status(http.StatusOK)
			// error會用grpc-status 及 grpc-status-details-bin 回傳
			wrappedGrpc.ServeHTTP(ctx.Writer, ctx.Request)
			return
		}

		ctx.Next()
	}
}
