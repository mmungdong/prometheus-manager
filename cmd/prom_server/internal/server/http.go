package server

import (
	"fmt"
	"time"

	alarmHistory "prometheus-manager/cmd/prom_server/internal/service/alarm_history"
	alarmPage "prometheus-manager/cmd/prom_server/internal/service/alarm_page"
	promDict "prometheus-manager/cmd/prom_server/internal/service/prom_dict"
	"prometheus-manager/cmd/prom_server/internal/service/strategy"
	strategyGroup "prometheus-manager/cmd/prom_server/internal/service/strategy_group"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"

	"prometheus-manager/cmd/prom_server/internal/conf"
	"prometheus-manager/cmd/prom_server/internal/service"
)

func NewHttpServer(server *conf.Server) *ginplus.GinEngine {
	middle := ginplus.NewMiddleware()
	var r *gin.Engine
	// 初始化gin实例
	if server.Mode == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(
		middle.Cors(),
		middle.Logger(server.Name, time.DateTime),
	)

	// 初始化ginplus实例
	ginplusEngine := ginplus.New(r,
		ginplus.WithAddr(fmt.Sprintf("%s:%d", server.Http.Host, server.Http.Port)),
		ginplus.AppendHttpMethodPrefixes(httpMethodPrefixes...),
		// 注册api模块
		ginplus.WithControllers(
			service.NewApi(
				service.WithStrategyApi(strategy.NewStrategy()),
				service.WithStrategyGroupApi(strategyGroup.NewStrategyGroup()),
				service.WithPromDictApi(promDict.NewPromDict()),
				service.WithAlarmPageApi(alarmPage.NewAlarmPage()),
				service.WithAlarmHistoryApi(alarmHistory.NewAlarmHistory()),
			),
		),
	)
	// 注册默认路由
	ginplusEngine.RegisterPing().RegisterSwaggerUI().RegisterMetrics()
	return ginplusEngine
}

var httpMethodPrefixes = []ginplus.HttpMethod{
	{
		Prefix: "Create",
		Method: ginplus.Post,
	}, {
		Prefix: "Update",
		Method: ginplus.Put,
	}, {
		Prefix: "Edit",
		Method: ginplus.Put,
	}, {
		Prefix: "Delete",
		Method: ginplus.Delete,
	}, {
		Prefix: "Detail",
		Method: ginplus.Get,
	}, {
		Prefix: "List",
		Method: ginplus.Get,
	}, {
		Prefix: "Login",
		Method: ginplus.Post,
	}, {
		Prefix: "Logout",
		Method: ginplus.Post,
	}, {
		Prefix: "Register",
		Method: ginplus.Post,
	},
}
