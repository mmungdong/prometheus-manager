package service

import (
	promDict "prometheus-manager/cmd/prom_server/internal/service/prom_dict"
	"prometheus-manager/cmd/prom_server/internal/service/strategy"
)

func WithStrategyApi(strategy *strategy.Strategy) ApiOption {
	return func(api *Api) {
		api.Strategy = strategy
	}
}

func WithPromDictApi(promDict *promDict.PromDict) ApiOption {
	return func(api *Api) {
		api.PromDict = promDict
	}
}
