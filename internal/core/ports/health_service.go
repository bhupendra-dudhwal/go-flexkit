package ports

type IHealth interface {
	Liveness(ctx IHandlerContext)
	Readiness(ctx IHandlerContext)
}
