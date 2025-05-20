package handler

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"
)

func NewHandler[T any](factory func() (ports.IHandlers, T)) (ports.IHandlers, T) {
	return factory()
}
