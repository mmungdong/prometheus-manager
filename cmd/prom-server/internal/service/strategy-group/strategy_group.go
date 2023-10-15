package StrategyGroup_group

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IStrategyGroup = (*StrategyGroup)(nil)

type (
	// IStrategyGroup ...
	IStrategyGroup interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// StrategyGroup ...
	StrategyGroup struct {
		// add child module
	}

	// Option ...
	Option func(*StrategyGroup)
)

// defaultStrategyGroup ...
func defaultStrategyGroup() *StrategyGroup {
	return &StrategyGroup{
		// add child module
	}
}

// NewStrategyGroup ...
func NewStrategyGroup(opts ...Option) *StrategyGroup {
	m := defaultStrategyGroup()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares 添加StrategyGroup模块的公共中间件
func (l *StrategyGroup) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethoderMiddlewares ...
func (l *StrategyGroup) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
