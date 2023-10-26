package dataStrategyGroup

import (
	query "github.com/aide-cloud/gorm-normalize"

	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"
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
	strategyGroup := &StrategyGroup{
		IAction:                  query.NewAction(query.WithDB[model.PromGroup](conn.GetMysqlDB())),
		PreloadCategoriesKey:     "Categories",
		PreloadPromStrategiesKey: "PromStrategies",
	}
	for _, o := range opts {
		o(strategyGroup)
	}

	return strategyGroup
}
