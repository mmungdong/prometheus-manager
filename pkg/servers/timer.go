package servers

import (
	"context"
	"time"

	"prometheus-manager/pkg/helper"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/go-kratos/kratos/v2/log"
)

var _ ginplus.Server = (*Timer)(nil)

type TimerCall func(ctx context.Context)

type Timer struct {
	call   TimerCall
	ticker *time.Ticker
	stop   chan struct{}
	logger *log.Helper
}

func (l *Timer) Start() error {
	l.logger.Info("[Timer] server starting")
	// 根据ticker的时间间隔，定时执行call
	go func() {
		defer helper.Recover("[Timer] server")
		for {
			select {
			case <-l.ticker.C:
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				defer cancel()
				l.call(ctx)
			case <-l.stop:
				return
			}
		}
	}()

	return nil
}

func (l *Timer) Stop() {
	l.logger.Info("[Timer] server stopping")
	l.stop <- struct{}{}
	l.ticker.Stop()
	l.logger.Info("[Timer] server stopped")
}

// NewTimer 创建一个定时器
func NewTimer(call TimerCall, ticker *time.Ticker, logger log.Logger) *Timer {
	return &Timer{
		call:   call,
		ticker: ticker,
		stop:   make(chan struct{}, 1),
		logger: log.NewHelper(log.With(logger, "module", "server/timer")),
	}
}
