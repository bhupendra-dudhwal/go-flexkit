package ports

type IAuth interface {
	Signin(ctx IHandlerContext)
	Signup(ctx IHandlerContext)
}
