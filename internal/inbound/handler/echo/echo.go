package echo

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/adapters"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"

	echov4 "github.com/labstack/echo/v4"
)

type echoHandler struct {
	engine *echov4.Echo
}

func NewEcho() (ports.IHandlers, *echov4.Echo) {
	e := echov4.New()
	return &echoHandler{engine: e}, e
}

func (e *echoHandler) SetHealthHandler(healthService ports.IHealth) {
	v1Group := e.engine.Group("/v1")
	{
		healthV1Group := v1Group.Group("/healthz")
		{
			healthV1Group.GET("/liveness", func(ctc echov4.Context) error {
				handlerContext := adapters.EchoAdapter{C: ctc}
				healthService.Liveness(&handlerContext)
				return nil
			})

			healthV1Group.GET("/readiness", func(ctc echov4.Context) error {
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
			authV1Group.POST("", func(ctc echov4.Context) error {
				handlerContext := adapters.EchoAdapter{C: ctc}
				authService.Signin(&handlerContext)
				return nil
			})
		}
	}
}

func (e *echoHandler) SetUserHandelr() {

}
