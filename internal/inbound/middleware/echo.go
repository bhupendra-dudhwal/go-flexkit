package middleware

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/constants"
	"github.com/labstack/echo/v4"
)

type EchoMiddleware struct{}

func NewEchoMiddleware() *EchoMiddleware {
	return &EchoMiddleware{}
}

func (e *EchoMiddleware) AddRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get(constants.RequestIDHeader.String())
			if requestID == "" {
				requestID = getRequestID()
			}

			// Set request id into context
			c.Set(constants.RequestIDKey.String(), requestID)

			// Set request id into response header
			c.Response().Header().Set(constants.RequestIDHeader.String(), requestID)
			return next(c)
		}
	}
}
