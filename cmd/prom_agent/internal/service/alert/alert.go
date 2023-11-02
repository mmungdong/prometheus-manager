package alert

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IAlert = (*Alert)(nil)

type (
	// IAlert ...
	IAlert interface {
		ginplus.IMiddleware
		ginplus.MethodeMiddleware
		// add interface method
	}

	// Alert ...
	Alert struct {
		// add child module
	}

	// Option ...
	Option func(*Alert)
)

// defaultAlert ...
func defaultAlert() *Alert {
	return &Alert{
		// add child module
	}
}

// NewAlert ...
func NewAlert(opts ...Option) *Alert {
	m := defaultAlert()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *Alert) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethodeMiddlewares ...
func (l *Alert) MethodeMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
