package service

import (
	alarmHistory "prometheus-manager/cmd/prom_server/internal/service/alarm_history"
	alarmPage "prometheus-manager/cmd/prom_server/internal/service/alarm_page"
	promDict "prometheus-manager/cmd/prom_server/internal/service/prom_dict"
	"prometheus-manager/cmd/prom_server/internal/service/strategy"
	strategyGroup "prometheus-manager/cmd/prom_server/internal/service/strategy_group"
)

func WithStrategyApi(strategy *strategy.Strategy) ApiOption {
	return func(api *Api) {
		api.Strategy = strategy
	}
}

func WithStrategyGroupApi(strategyGroup *strategyGroup.StrategyGroup) ApiOption {
	return func(api *Api) {
		api.StrategyGroup = strategyGroup
	}
}

func WithPromDictApi(promDict *promDict.PromDict) ApiOption {
	return func(api *Api) {
		api.PromDict = promDict
	}
}

func WithAlarmPageApi(page *alarmPage.AlarmPage) ApiOption {
	return func(api *Api) {
		api.AlarmPage = page
	}
}

func WithAlarmHistoryApi(history *alarmHistory.AlarmHistory) ApiOption {
	return func(api *Api) {
		api.History = history
	}
}
