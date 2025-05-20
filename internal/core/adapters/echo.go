package adapters

import (
	echov4 "github.com/labstack/echo/v4"
)

type EchoAdapter struct {
	C echov4.Context
}

func (h *EchoAdapter) Bind(obj any) error {
	return h.C.Bind(obj)
}

func (h *EchoAdapter) JSON(code int, obj any) {
	h.C.JSON(code, obj)
}

func (h *EchoAdapter) XML(code int, obj any) {
	h.C.XML(code, obj)
}

func (h *EchoAdapter) BindJSON(obj any) error {
	return h.Bind(obj)
}

func (h *EchoAdapter) BindXML(obj any) error {
	return h.Bind(obj)
}

func (h *EchoAdapter) Param(key string) string {
	return h.C.Param(key)
}

func (h *EchoAdapter) Query(key string) string {
	return h.C.QueryParams().Get(key)
}

func (h *EchoAdapter) QueryParam(key string) string {
	return h.Query(key)
}
