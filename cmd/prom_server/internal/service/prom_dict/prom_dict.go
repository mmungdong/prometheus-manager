package promDict

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IPromDict = (*PromDict)(nil)

type (
	// IPromDict ...
	IPromDict interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// PromDict ...
	PromDict struct {
		// add child module
	}

	// PromDictOption ...
	PromDictOption func(*PromDict)
)

// defaultPromDict ...
func defaultPromDict() *PromDict {
	return &PromDict{
		// add child module
	}
}

// NewPromDict ...
func NewPromDict(opts ...PromDictOption) *PromDict {
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

// MethoderMiddlewares ...
func (l *PromDict) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}