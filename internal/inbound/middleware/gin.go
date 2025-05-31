package middleware

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/constants"
	"github.com/gin-gonic/gin"
)

type GinMiddleware struct{}

func NewGinMiddleware() *GinMiddleware {
	return &GinMiddleware{}
}

func (g *GinMiddleware) AddRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(constants.RequestIDHeader.String())
		if requestID == "" {
			requestID = getRequestID()
		}

		// Set request id in context
		c.Set(constants.RequestIDKey.String(), requestID)

		// Set request id in header
		c.Writer.Header().Set(constants.RequestIDHeader.String(), requestID)

		c.Next()
	}
}
