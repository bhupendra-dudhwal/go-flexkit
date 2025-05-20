package main

import (
	"github.com/bhupendra-dudhwal/go-flexkit/internal/builder"
)

func main() {
	builder.NewBuilder().
		SetConfig().
		SetAuthService().
		SetHandler().
		SetServer().
		BuildAndStart()
}
