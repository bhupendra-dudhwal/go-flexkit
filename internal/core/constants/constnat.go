package constants

type Handler string

const (
	ECHO     Handler = "ECHO"
	GIN      Handler = "GIN"
	FASTHTTP Handler = "FASTHTTP"
	MUX      Handler = "MUX"
)
