package service

import (
	alarmHistory "prometheus-manager/cmd/prom_server/internal/service/alarm_history"
	alarmPage "prometheus-manager/cmd/prom_server/internal/service/alarm_page"
	promDict "prometheus-manager/cmd/prom_server/internal/service/prom_dict"
	"prometheus-manager/cmd/prom_server/internal/service/strategy"
	strategyGroup "prometheus-manager/cmd/prom_server/internal/service/strategy_group"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var _ IApi = (*Api)(nil)

type (
	// IApi ...
	IApi interface {
		ginplus.Middlewarer
		ginplus.MethoderMiddlewarer
		// add interface method
	}

	// Api ...
	Api struct {
		// add child module
		Strategy      *strategy.Strategy
		StrategyGroup *strategyGroup.StrategyGroup
		PromDict      *promDict.PromDict
		Page          *alarmPage.AlarmPage
		History       *alarmHistory.AlarmHistory
	}

	// ApiOption ...
	ApiOption func(*Api)
)

// defaultApi ...
func defaultApi() *Api {
	return &Api{
		// add child module
	}
}

// NewApi ...
func NewApi(opts ...ApiOption) *Api {
	m := defaultApi()
	for _, o := range opts {
		o(m)
	}

	return m
}

// Middlewares 添加Api模块的公共中间件
func (l *Api) Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			// your code ...
		},
		// other ...
	}
}

// MethoderMiddlewares ...
func (l *Api) MethoderMiddlewares() map[string][]gin.HandlerFunc {
	return map[string][]gin.HandlerFunc{
		// your method middlewares
	}
}
