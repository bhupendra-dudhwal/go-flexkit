package service

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"
)

type authService struct{}

func NewAuth() ports.IAuth {
	return &authService{}
}

func (a *authService) Signin(ctx ports.IHandlerContext) {

}

func (a *authService) Signup(ctx ports.IHandlerContext) {

}
