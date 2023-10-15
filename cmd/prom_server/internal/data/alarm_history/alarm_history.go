package alarm_history

import (
	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// AlarmHistory mapped from table <alarm_history>
	AlarmHistory struct {
		query.IAction[model.PromAlarmHistory]
	}

	// AlarmHistoryOption AlarmHistory's option
	AlarmHistoryOption func(*AlarmHistory)
)

// NewAlarmHistory new a AlarmHistory
func NewAlarmHistory(options ...AlarmHistoryOption) *AlarmHistory {
	alarmHistory := &AlarmHistory{
		IAction: query.NewAction(query.WithDB[model.PromAlarmHistory](conn.GetMysqlDB())),
	}
	for _, option := range options {
		option(alarmHistory)
	}
	return alarmHistory
}
