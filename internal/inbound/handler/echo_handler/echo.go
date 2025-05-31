package echohandler

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/adapters"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/inbound/middleware"

	"github.com/labstack/echo/v4"
)

type echoHandler struct {
	engine     *echo.Echo
	middleware *middleware.EchoMiddleware
}

func NewEcho() (ports.IHandlers, *echo.Echo) {
	middleware := middleware.NewEchoMiddleware() // Get echo middleware
	e := echo.New()
	e.Use(middleware.AddRequestID()) // Global middleware that will be used for all requests
	return &echoHandler{
		engine:     e,
		middleware: middleware,
	}, e
}

func (e *echoHandler) SetHealthHandler(healthService ports.IHealth) {
	v1Group := e.engine.Group("/v1")
	{
		healthV1Group := v1Group.Group("/healthz")
		{
			healthV1Group.GET("/liveness", func(ctc echo.Context) error {
				handlerContext := adapters.EchoAdapter{C: ctc}
				healthService.Liveness(&handlerContext)
				return nil
			})

			healthV1Group.GET("/readiness", func(ctc echo.Context) error {
				handlerContext := adapters.EchoAdapter{C: ctc}
				healthService.Readiness(&handlerContext)
				return nil
			})
		}
	}
}

func (e *echoHandler) SetAuthHandler(authService ports.IAuth) {
	v1Group := e.engine.Group("/v1")
	{
		authV1Group := v1Group.Group("/auth")
		{
			authV1Group.POST("", func(ctc echo.Context) error {
				handlerContext := adapters.EchoAdapter{C: ctc}
				authService.Signin(&handlerContext)
				return nil
			})
		}
	}
}

func (e *echoHandler) SetUserHandelr() {

}
