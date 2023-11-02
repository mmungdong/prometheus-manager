package promDict

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IPromDict = (*PromDict)(nil)

type (
	// IPromDict ...
	IPromDict interface {
		ginplus.IMiddleware
		ginplus.MethodeMiddleware
		// add interface method
	}

	// PromDict ...
	PromDict struct {
		// add child module
	}

	// Option ...
	Option func(*PromDict)
)

// defaultPromDict ...
func defaultPromDict() *PromDict {
	return &PromDict{
		// add child module
	}
}

// NewPromDict ...
func NewPromDict(opts ...Option) *PromDict {
	m := defaultPromDict()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *PromDict) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethodeMiddlewares ...
func (l *PromDict) MethodeMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
