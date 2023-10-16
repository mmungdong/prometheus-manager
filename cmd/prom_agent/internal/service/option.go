package service

import "prometheus-manager/cmd/prom_agent/internal/service/alert"

func WithAlert(a *alert.Alert) ApiOption {
	return func(api *Api) {
		api.Alert = a
	}
}
