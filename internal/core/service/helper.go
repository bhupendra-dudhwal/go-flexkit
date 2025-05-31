package service

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/constants"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"
)

func getRequestID(ctx ports.IHandlerContext) string {
	if val, ok := ctx.GetContext(constants.RequestIDKey.String()); ok {
		if reqID, ok := val.(string); ok {
			return reqID
		}
	}
	return ""
}
