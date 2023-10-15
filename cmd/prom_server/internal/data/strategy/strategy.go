package strategy

import (
	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// Strategy ...
	Strategy struct {
		query.IAction[model.PromStrategy]
	}

	// StrategyOption ...
	StrategyOption func(*Strategy)
)

// NewStrategy new a Strategy instance
func NewStrategy(opts ...StrategyOption) *Strategy {
	strategy := &Strategy{
		IAction: query.NewAction(query.WithDB[model.PromStrategy](conn.GetMysqlDB())),
	}
	for _, o := range opts {
		o(strategy)
	}

	return strategy
}
