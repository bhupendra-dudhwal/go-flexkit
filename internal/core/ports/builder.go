package ports

type IBuilder interface {
	SetConfig() IBuilder
	SetAuthService() IBuilder
	SetHandler() IBuilder
	SetServer() IBuilder
	BuildAndStart()
}
