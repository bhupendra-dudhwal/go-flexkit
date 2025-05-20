package adapters

import (
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	C *gin.Context
}

func (h *GinAdapter) Bind(obj any) error {
	return h.C.Bind(obj)
}

func (h *GinAdapter) JSON(code int, obj any) {
	h.C.JSON(code, obj)
}

func (h *GinAdapter) XML(code int, obj any) {
	h.C.XML(code, obj)
}

func (h *GinAdapter) BindJSON(obj any) error {
	return h.C.BindJSON(obj)
}

func (h *GinAdapter) BindXML(obj any) error {
	return h.C.BindXML(obj)
}

func (h *GinAdapter) Param(key string) string {
	return h.C.Param(key)
}

func (h *GinAdapter) Query(key string) string {
	return h.C.Query(key)
}
