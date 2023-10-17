package server

import (
	"fmt"

	"prometheus-manager/cmd/prom_agent/internal/conf"
	"prometheus-manager/cmd/prom_agent/internal/service"
	"prometheus-manager/cmd/prom_agent/internal/service/alert"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

func NewHttpServer(server *conf.Server, looger log.Logger) *ginplus.GinEngine {
	logHelper := log.NewHelper(log.With(looger, "module", "server/server"))
	logHelper.Infof("HttpServer starting")
	// middle := ginplus.NewMiddleware()
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
	// middle.Cors(),
	// middle.Logger(server.Name, time.DateTime),
	)

	// 初始化ginplus实例
	ginplusEngine := ginplus.New(r,
		ginplus.WithAddr(fmt.Sprintf("%s:%d", server.Http.Host, server.Http.Port)),
		ginplus.AppendHttpMethodPrefixes(httpMethodPrefixes...),
		// 注册api模块
		ginplus.WithControllers(
			service.NewApi(
				service.WithAlert(alert.NewAlert()),
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
