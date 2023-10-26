package dataAlarmPage

import (
	query "github.com/aide-cloud/gorm-normalize"

	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"
)

type (
	// AlarmPage ...
	AlarmPage struct {
		query.IAction[model.PromAlarmPage]
	}

	// AlarmPageOption ...
	AlarmPageOption func(*AlarmPage)
)

// NewAlarmPage new a AlarmPage instance
func NewAlarmPage(opts ...AlarmPageOption) *AlarmPage {
	page := &AlarmPage{
		IAction: query.NewAction(query.WithDB[model.PromAlarmPage](conn.GetMysqlDB())),
	}
	for _, o := range opts {
		o(page)
	}
	return page
}
