package ports

type IHandlerContext interface {
	Bind(any) error
	BindJSON(any) error
	BindXML(any) error
	JSON(int, any)
	XML(int, any)
	Param(string) string
	Query(string) string
}

type IHandlers interface {
	SetHealthHandler(IHealth)
	SetAuthHandler(IAuth)
	SetUserHandelr()
}
