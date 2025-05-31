package ports

type IHandlerContext interface {
	Bind(any) error
	BindJSON(any) error
	BindXML(any) error
	JSON(int, any)
	XML(int, any)
	Param(string) string
	Query(string) string
	GetContext(key string) (value any, exists bool)
}

type IHandlers interface {
	SetHealthHandler(IHealth)
	SetAuthHandler(IAuth)
	SetUserHandelr()
}

type IHandlerMiddlewares interface {
	RequestID()
}
