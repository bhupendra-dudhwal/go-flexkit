# go-flexkit

A flexible and extensible Go starter kit for building HTTP services using Clean Architecture principles.  
Supports multiple web frameworks like **Gin**, **Echo**, and **Mux** out of the box.

## 🌟 Features

- 🔁 Pluggable HTTP frameworks: Gin, Echo, Mux
- 🧼 Clean Architecture (Ports & Adapters)
- ✅ Framework-agnostic service layer
- 🧪 Easily testable business logic
- 🔌 Adapter layer for unified request context handling
- ⚡ Fast setup for new services

## 📁 Project Structure

```
├── cmd
│   └── http
│       └── server.go
├── config
│   └── config.yaml
├── go.mod
├── go.sum
└── internal
    ├── builder
    │   └── builder.go
    ├── core
    │   ├── adapters
    │   │   ├── echo.go
    │   │   └── gin.go
    │   ├── constants
    │   │   ├── constnat.go
    │   │   └── helper.go
    │   ├── model
    │   │   ├── config.go
    │   │   └── response.go
    │   ├── ports
    │   │   ├── auth_service.go
    │   │   ├── builder.go
    │   │   ├── handler.go
    │   │   └── health_service.go
    │   ├── repository
    │   └── service
    │       ├── auth.go
    │       └── health.go
    ├── inbound
    │   ├── handler
    │   │   ├── echo
    │   │   │   └── echo.go
    │   │   ├── gin
    │   │   │   └── gin.go
    │   │   └── handler.go
    │   ├── middleware
    │   └── response
    ├── outbound
    │   ├── api
    │   └── db
    └── utils
