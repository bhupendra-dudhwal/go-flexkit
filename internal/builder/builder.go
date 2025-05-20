package builder

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/constants"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/model"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/ports"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/core/service"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/inbound/handler"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/inbound/handler/echo"
	"github.com/bhupendra-dudhwal/go-flexkit/internal/inbound/handler/gin"

	ginEngine "github.com/gin-gonic/gin"
	echov4 "github.com/labstack/echo/v4"

	"gopkg.in/yaml.v3"
)

type builder struct {
	config        *model.Config
	handlerPorts  ports.IHandlers
	authService   ports.IAuth
	healthService ports.IHealth
	server        *http.Server
}

func NewBuilder() ports.IBuilder {
	return &builder{
		config: &model.Config{},
		server: &http.Server{},
	}
}

func (b *builder) SetConfig() ports.IBuilder {
	configBytes, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	cnf := model.Config{}
	if err := yaml.Unmarshal(configBytes, &cnf); err != nil {
		log.Fatalf("error: %v", err)
	}
	b.config = &cnf

	return b
}

func (b *builder) SetAuthService() ports.IBuilder {
	b.authService = service.NewAuth()
	b.healthService = service.NewHealthService()
	return b
}

func (b *builder) SetHandler() ports.IBuilder {

	switch b.config.Handler.Handler {
	case constants.GIN:
		b.handlerPorts, b.server.Handler = handler.NewHandler[*ginEngine.Engine](gin.NewGin)
	case constants.ECHO:
		b.handlerPorts, b.server.Handler = handler.NewHandler[*echov4.Echo](echo.NewEcho)
	default:
		log.Fatalf("%s handler framework do not support", b.config.Handler.Handler)
	}

	b.handlerPorts.SetHealthHandler(b.healthService)
	b.handlerPorts.SetAuthHandler(b.authService)

	return b
}

func (b *builder) SetServer() ports.IBuilder {
	b.server.Addr = fmt.Sprintf(":%d", b.config.App.Port)
	return b
}

func (b *builder) BuildAndStart() {

	go func() {
		log.Printf("server starting on :%d", b.config.App.Port)
		if err := b.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeoutCause(context.Background(), 10*time.Second, errors.New("server interrupt by os signal"))
	defer cancel()

	if err := b.server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	} else {
		log.Println("server gracefully stopped")
	}
}
