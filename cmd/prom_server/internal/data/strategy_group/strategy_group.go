package dataStrategyGroup

import (
	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// StrategyGroup ...
	StrategyGroup struct {
		query.IAction[model.PromGroup]

		PreloadCategoriesKey     query.AssociationKey
		PreloadPromStrategiesKey query.AssociationKey
	}

	// StrategyGroupOption ...
	StrategyGroupOption func(*StrategyGroup)
)

// NewStrategyGroup new a StrategyGroup instance
func NewStrategyGroup(opts ...StrategyGroupOption) *StrategyGroup {
	strategy_group := &StrategyGroup{
		IAction:                  query.NewAction(query.WithDB[model.PromGroup](conn.GetMysqlDB())),
		PreloadCategoriesKey:     "Categories",
		PreloadPromStrategiesKey: "PromStrategies",
	}
	for _, o := range opts {
		o(strategy_group)
	}

	return strategy_group
}
