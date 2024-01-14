package v1

import (
	"estore/internal/fs"
	"estore/internal/mariadb"
	"estore/pkg/module/logger"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jigadhirasu/rex"
	"go.uber.org/zap"
)

func Http(router *gin.Engine) {
	g := router.Group("/api/v1")

	hs := reflect.ValueOf(HttpService{})

	for i := 0; i < hs.NumMethod(); i++ {
		method := hs.Method(i)
		method.Call([]reflect.Value{reflect.ValueOf(g)})
	}
}

type HttpService struct {
}

func (HttpService) Static(router gin.IRouter) {

	dbname := "test"

	router.GET("/static/:id", func(gtx *gin.Context) {

		id := gtx.Param("id")

		pipe := rex.Pipe3(
			rex.Tap(mariadb.F0GetDB[string](dbname)),
			rex.Map(fs.F1FindFilestore(dbname)),
			rex.Tap(func(ctx rex.Context, filename string) error {
				logger.Out().Info("filepath",
					zap.String("id", id),
					zap.String("fullname", filename),
				)

				http.ServeFile(gtx.Writer, gtx.Request, filename)
				return nil
			}),
		)(rex.From(id))

		ctx := rex.NewContext(gtx)

		item := <-pipe(ctx)()

		if _, err := item(); err != nil {
			gtx.AbortWithError(http.StatusInternalServerError, err)
		}
	})

}
