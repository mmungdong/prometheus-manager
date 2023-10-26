package alarmPage

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IAlarmPage = (*AlarmPage)(nil)

type (
	// IAlarmPage ...
	IAlarmPage interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// AlarmPage ...
	AlarmPage struct {
		// add child module
	}

	// Option ...
	Option func(*AlarmPage)
)

// defaultAlarmPage ...
func defaultAlarmPage() *AlarmPage {
	return &AlarmPage{
		// add child module
	}
}

// NewAlarmPage ...
func NewAlarmPage(opts ...Option) *AlarmPage {
	m := defaultAlarmPage()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares ...
func (l *AlarmPage) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethoderMiddlewares ...
func (l *AlarmPage) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
