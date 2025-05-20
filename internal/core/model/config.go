package model

import "github.com/bhupendra-dudhwal/go-flexkit/internal/core/constants"

type AppObj struct {
}

type Config struct {
	App     *App     `yaml:"app"`
	Handler *Handler `yaml:"handler"`
}

type App struct {
	Port        int    `yaml:"port"`
	Environment string `yaml:"environment"`
}

type Handler struct {
	Handler constants.Handler `yaml:"handler"`
}
