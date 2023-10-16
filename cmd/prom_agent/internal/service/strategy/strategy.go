package strategy

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IStrategy = (*Strategy)(nil)

type (
	// IStrategy ...
	IStrategy interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// Strategy ...
	Strategy struct {
		// add child module
	}

	// StrategyOption ...
	StrategyOption func(*Strategy)
)

// defaultStrategy ...
func defaultStrategy() *Strategy {
	return &Strategy{
		// add child module
	}
}

// NewStrategy ...
func NewStrategy(opts ...StrategyOption) *Strategy {
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

// MethoderMiddlewares ...
func (l *Strategy) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
