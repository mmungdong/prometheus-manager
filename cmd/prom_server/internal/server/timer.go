package server

import (
	"context"
	"time"

	"prometheus-manager/cmd/prom_server/internal/conf"
	"prometheus-manager/cmd/prom_server/internal/service/strategy/publish"
	"prometheus-manager/pkg/servers"

	"github.com/go-kratos/kratos/v2/log"
)

func NewTimer(pushStrategy *conf.PushStrategy, logger log.Logger) *servers.Timer {
	interval := time.Duration(pushStrategy.GetIntervel()) * time.Second
	ticker := time.NewTicker(interval)
	loggerHelper := log.NewHelper(log.With(logger, "module", "server/Timer"))

	publishService := publish.NewPublish()

	var count int64
	call := func(ctx context.Context) {
		if !pushStrategy.GetEnable() {
			return
		}
		count++
		loggerHelper.Info("TimerCallFunc: ", count)

		pushed, err := publishService.PostStrategy(ctx, &publish.StrategyReq{})
		if err != nil {
			loggerHelper.Errorf("[Timer] call error: %v", err)
			return
		}

		log.Info("pushed: ", pushed)
	}
	return servers.NewTimer(call, ticker, logger)
}
