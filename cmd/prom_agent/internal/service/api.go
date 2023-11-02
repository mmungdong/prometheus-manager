package service

import (
	"prometheus-manager/cmd/prom_agent/internal/service/alert"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IApi = (*Api)(nil)

type (
	// IApi ...
	IApi interface {
		ginplus.IMiddleware
		ginplus.MethodeMiddleware
		// add interface method
	}

	// Api ...
	Api struct {
		Alert *alert.Alert
	}

	// ApiOption ...
	ApiOption func(*Api)
)

// defaultApi ...
func defaultApi() *Api {
	return &Api{
		// add child module
	}
}

// NewApi ...
func NewApi(opts ...ApiOption) *Api {
	m := defaultApi()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *Api) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethodeMiddlewares ...
func (l *Api) MethodeMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
