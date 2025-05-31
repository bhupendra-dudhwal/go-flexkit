package ginhandler

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/adapters"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/inbound/middleware"
	"github.com/gin-gonic/gin"
)

type ginHandler struct {
	engine     *gin.Engine
	middleware *middleware.GinMiddleware
}

func NewGin() (ports.IHandlers, *gin.Engine) {
	middleware := middleware.NewGinMiddleware() // Get gin middleware
	g := gin.New()
	g.Use(middleware.AddRequestID()) // Global middleware that will be used for all requests
	return &ginHandler{
		engine:     g,
		middleware: middleware,
	}, g
}

func (g *ginHandler) SetHealthHandler(healthService ports.IHealth) {
	v1Group := g.engine.Group("/v1")
	{
		healthV1Group := v1Group.Group("/healthz")
		{
			healthV1Group.GET("/liveness", func(ctx *gin.Context) {
				handlerContext := adapters.GinAdapter{C: ctx}
				healthService.Liveness(&handlerContext)
			})

			healthV1Group.GET("/readiness", func(ctc *gin.Context) {
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
			authV1Group.POST("", func(ctx *gin.Context) {
				handlerContext := adapters.GinAdapter{C: ctx}
				authService.Signin(&handlerContext)
			})
		}
	}
}

func (g *ginHandler) SetUserHandelr() {

}
