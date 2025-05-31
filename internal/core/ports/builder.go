package ports

type IBuilder interface {
	SetConfig() IBuilder
	SetServices() IBuilder
	SetHandler() IBuilder
	SetServer() IBuilder
	BuildAndStart()
}
