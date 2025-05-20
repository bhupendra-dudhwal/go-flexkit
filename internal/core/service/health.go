package service

import (
	"net/http"

	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/model"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"
)

type heathService struct{}

func NewHealthService() ports.IHealth {
	return &heathService{}
}

func (h *heathService) Liveness(ctx ports.IHandlerContext) {
	ctx.XML(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "App is live",
	})
}

func (h *heathService) Readiness(ctx ports.IHandlerContext) {
	ctx.JSON(http.StatusOK, model.Response{
		Code:    http.StatusOK,
		Status:  true,
		Message: "App is ready to serve the request",
	})
}
