package gin

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/adapters"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"

	ginEngine "github.com/gin-gonic/gin"
)

type ginHandler struct {
	engine *ginEngine.Engine
}

func NewGin() (ports.IHandlers, *ginEngine.Engine) {
	g := ginEngine.New()
	return &ginHandler{engine: g}, g
}

func (g *ginHandler) SetHealthHandler(healthService ports.IHealth) {
	v1Group := g.engine.Group("/v1")
	{
		healthV1Group := v1Group.Group("/healthz")
		{
			healthV1Group.GET("/liveness", func(ctx *ginEngine.Context) {
				handlerContext := adapters.GinAdapter{C: ctx}
				healthService.Liveness(&handlerContext)
			})

			healthV1Group.GET("/readiness", func(ctc *ginEngine.Context) {
				handlerContext := adapters.GinAdapter{C: ctc}
				healthService.Readiness(&handlerContext)
			})
		}
	}
}

func (g *ginHandler) SetAuthHandler(authService ports.IAuth) {
	v1Group := g.engine.Group("/v1")
	{
		authV1Group := v1Group.Group("/auth")
		{
			authV1Group.POST("", func(ctx *ginEngine.Context) {
				handlerContext := adapters.GinAdapter{C: ctx}
				authService.Signin(&handlerContext)
			})
		}
	}
}

func (g *ginHandler) SetUserHandelr() {

}
