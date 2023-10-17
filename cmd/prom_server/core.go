package main

import (
	"prometheus-manager/cmd/prom_server/internal/conf"
	"prometheus-manager/pkg/conn"

	ginplus "github.com/aide-cloud/gin-plus"
)

func Init() *conf.Bootstrap {
	bc := conf.NewBootstrap(conf.WithConfigFile(*configPath))
	if bc == nil {
		panic("conf.NewBootstrap() error")
	}

	if bc.Server != nil && bc.Server.Name != "" {
		ServiceName = bc.Server.Name
	}

	ginplus.Logger().Sugar().Infof("%s version: %s", ServiceName, Version)

	if bc.Data != nil {
		mysqlConf := bc.GetData().GetMysql()
		if mysqlConf != nil && mysqlConf.GetDSN() != "" {
			conn.InitMysqlDB(mysqlConf.GetDSN(), mysqlConf.GetDebug())
		}
	}

	return bc
}
