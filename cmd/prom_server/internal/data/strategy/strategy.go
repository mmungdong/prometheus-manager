package dataStrategy

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/gorm"

	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"
)

type (
	// Strategy ...
	Strategy struct {
		query.IAction[model.PromStrategy]

		PreloadAlarmPagesKey query.AssociationKey
		PreloadCategoriesKey query.AssociationKey
		PreloadAlertLevelKey query.AssociationKey
		PreloadGroupInfoKey  query.AssociationKey
	}

	// StrategyOption ...
	StrategyOption func(*Strategy)
)

// NewStrategy new a Strategy instance
func NewStrategy(opts ...StrategyOption) *Strategy {
	strategy := &Strategy{
		IAction: query.NewAction(query.WithDB[model.PromStrategy](conn.GetMysqlDB())),

		PreloadAlarmPagesKey: "AlarmPages",
		PreloadCategoriesKey: "Categories",
		PreloadAlertLevelKey: "AlertLevel",
		PreloadGroupInfoKey:  "GroupInfo",
	}
	for _, o := range opts {
		o(strategy)
	}

	return strategy
}

// PreloadAlarmPages preload alarm_pages
func (l *Strategy) PreloadAlarmPages() query.Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(string(l.PreloadAlarmPagesKey))
	}
}

// PreloadCategories preload categories
func (l *Strategy) PreloadCategories() query.Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(string(l.PreloadCategoriesKey))
	}
}

// PreloadAlertLevel preload alert_level
func (l *Strategy) PreloadAlertLevel() query.Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(string(l.PreloadAlertLevelKey))
	}
}

// PreloadGroupInfo preload group_info
func (l *Strategy) PreloadGroupInfo() query.Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(string(l.PreloadGroupInfoKey))
	}
}
