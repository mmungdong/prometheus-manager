package service

import (
	"prometheus-manager/cmd/prom-server/internal/service/strategy"
)

func WithStrategyApi(strategy *strategy.Strategy) ApiOption {
	return func(api *Api) {
		api.Strategy = strategy
	}
}
