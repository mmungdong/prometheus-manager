package strategy

import (
	"context"
	"mime/multipart"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
	"prometheus-manager/pkg/model"
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
		strategyRepository Repository
	}

	// Option ...
	Option func(*Strategy)

	// Repository ...
	Repository interface {
		// ParseStrategyFiles 解析策略文件
		ParseStrategyFiles(ctx context.Context, yamlFiles []*multipart.FileHeader) ([]*model.PromStrategy, error)
		// BatchCreateStrategy 批量创建策略
		BatchCreateStrategy(ctx context.Context, strategies []*model.PromStrategy) error
	}
)

// defaultStrategy ...
func defaultStrategy() *Strategy {
	return &Strategy{
		//strategyRepository: repository_impl.NewStrategyRepository(),
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

// Middlewares 添加Strategy模块的公共中间件
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

// WithStrategyRepository ...
func WithStrategyRepository(repository Repository) Option {
	return func(m *Strategy) {
		m.strategyRepository = repository
	}
}
