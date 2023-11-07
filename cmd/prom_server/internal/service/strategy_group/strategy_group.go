package strategyGroup

import (
	"context"
	"mime/multipart"
	"os"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/strategy"
)

var _ IStrategyGroup = (*StrategyGroup)(nil)

type (
	// IStrategyGroup ...
	IStrategyGroup interface {
		ginplus.IMiddleware
		ginplus.MethodeMiddleware
		// add interface method
	}

	// StrategyGroup ...
	StrategyGroup struct {
		// add child module
		groupRepository GroupRepository
	}

	// GroupRepository ...
	GroupRepository interface {
		GroupImportRepository
		GroupExportRepository
	}

	// GroupImportRepository ...
	GroupImportRepository interface {
		ParseStrategyGroupFile(ctx context.Context, yamlFiles []*multipart.FileHeader) ([]*model.PromGroup, error)
		BatchCreateStrategyGroup(ctx context.Context, groups []*model.PromGroup) error
	}

	// GroupExportRepository ...
	GroupExportRepository interface {
		GetStrategyGroupDetailById(ctx context.Context, id uint) (*model.PromGroup, error)
		ExportStrategyGroupFile(ctx context.Context, groupDetail *strategy.Group) (*os.File, error)
	}

	// Option ...
	Option func(*StrategyGroup)
)

// defaultStrategyGroup ...
func defaultStrategyGroup() *StrategyGroup {
	return &StrategyGroup{
		//groupRepository: repository_impl.NewGroupRepository(),
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

// Middlewares ...
func (l *StrategyGroup) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethodeMiddlewares ...
func (l *StrategyGroup) MethodeMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}

// WithGroupRepository ...
func WithGroupRepository(repository GroupRepository) Option {
	return func(m *StrategyGroup) {
		m.groupRepository = repository
	}
}
