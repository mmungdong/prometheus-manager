package publish

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IPublish = (*Publish)(nil)

type (
	// IPublish ...
	IPublish interface {
		ginplus.IMiddleware
		ginplus.MethodeMiddleware
		// add interface method
	}

	// Publish ...
	Publish struct {
		// add child module
	}

	// Option ...
	Option func(*Publish)
)

// defaultPublish ...
func defaultPublish() *Publish {
	return &Publish{
		// add child module
	}
}

// NewPublish ...
func NewPublish(opts ...Option) *Publish {
	m := defaultPublish()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *Publish) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethodeMiddlewares ...
func (l *Publish) MethodeMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
