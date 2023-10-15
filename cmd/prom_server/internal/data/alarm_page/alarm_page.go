package alarmPage

import (
	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
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
	AlarmPage := &AlarmPage{
		IAction: query.NewAction(query.WithDB[model.PromAlarmPage](conn.GetMysqlDB())),
	}
	for _, o := range opts {
		o(AlarmPage)
	}
	return AlarmPage
}
