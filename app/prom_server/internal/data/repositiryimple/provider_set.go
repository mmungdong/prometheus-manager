package repositiryimple

import (
	"github.com/google/wire"

	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmhistory"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/alarmpage"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/ping"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/promdict"
	"prometheus-manager/app/prom_server/internal/data/repositiryimple/strategy"
)

// ProviderSetRepository 注入repository依赖
var ProviderSetRepository = wire.NewSet(
	ping.NewPingRepo,
	promdict.NewPromDictRepo,
	strategy.NewStrategyRepo,
	alarmpage.NewAlarmPageRepo,
	alarmhistory.NewAlarmHistoryRepo,
)