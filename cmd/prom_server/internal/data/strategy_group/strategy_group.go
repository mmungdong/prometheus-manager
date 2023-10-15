package strategyGroup

import (
	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// StrategyGroup ...
	StrategyGroup struct {
		query.IAction[model.PromGroup]
	}

	// StrategyGroupOption ...
	StrategyGroupOption func(*StrategyGroup)
)

// NewStrategyGroup new a StrategyGroup instance
func NewStrategyGroup(opts ...StrategyGroupOption) *StrategyGroup {
	strategy_group := &StrategyGroup{
		IAction: query.NewAction(query.WithDB[model.PromGroup](conn.GetMysqlDB())),
	}
	for _, o := range opts {
		o(strategy_group)
	}

	return strategy_group
}
