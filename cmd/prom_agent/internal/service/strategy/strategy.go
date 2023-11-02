package strategy

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IStrategy = (*Strategy)(nil)

type (
	// IStrategy ...
	IStrategy interface {
		ginplus.IMiddleware
		ginplus.MethodeMiddleware
		// add interface method
	}

	// Strategy ...
	Strategy struct {
		// add child module
	}

	// Option ...
	Option func(*Strategy)
)

// defaultStrategy ...
func defaultStrategy() *Strategy {
	return &Strategy{
		// add child module
	}
}

// NewStrategy ...
func NewStrategy(opts ...Option) *Strategy {
	m := defaultStrategy()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *Strategy) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethodeMiddlewares ...
func (l *Strategy) MethodeMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
