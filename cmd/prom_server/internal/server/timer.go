package server

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"prometheus-manager/cmd/prom_server/internal/conf"
	"prometheus-manager/cmd/prom_server/internal/service/strategy/publish"
	"prometheus-manager/pkg/servers"
)

func NewTimeServer(bc *conf.Bootstrap, logger log.Logger) *servers.Timer {
	pushStrategy := bc.GetPushStrategy()
	interval := pushStrategy.GetInterval() * time.Second
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

		pushed, err := publishService.PostStrategy(ctx, &publish.StrategyReq{Topic: pushStrategy.GetTopic()})
		if err != nil {
			loggerHelper.Errorf("[Timer] call error: %v", err)
			return
		}

		log.Info("pushed: ", pushed)
	}
	return servers.NewTimer(call, ticker, logger)
}
