package alarmHistory

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IAlarmHistory = (*AlarmHistory)(nil)

type (
	// IAlarmHistory ...
	IAlarmHistory interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// AlarmHistory ...
	AlarmHistory struct {
		// add child module
	}

	// AlarmHistoryOption ...
	AlarmHistoryOption func(*AlarmHistory)
)

// defaultAlarmHistory ...
func defaultAlarmHistory() *AlarmHistory {
	return &AlarmHistory{
		// add child module
	}
}

// NewAlarmHistory ...
func NewAlarmHistory(opts ...AlarmHistoryOption) *AlarmHistory {
	m := defaultAlarmHistory()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *AlarmHistory) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethoderMiddlewares ...
func (l *AlarmHistory) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
