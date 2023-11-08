package main

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"prometheus-manager/cmd/prom_server/internal/conf"
	"prometheus-manager/pkg/conn"
)

func Init() *conf.Bootstrap {
	bc := conf.NewBootstrap(conf.WithConfigFile(*configPath))
	if bc == nil {
		panic("conf.NewBootstrap() error")
	}

	if bc.Server != nil && bc.Server.Name != "" {
		ServiceName = bc.Server.Name
	}

	if bc.Data != nil {
		mysqlConf := bc.GetData().GetMysql()
		if mysqlConf != nil && mysqlConf.GetDSN() != "" {
			conn.InitMysqlDB(mysqlConf.GetDSN(), mysqlConf.GetDebug())
		}
	}

	ginplus.Logger().Sugar().Infof("%s version: %s", ServiceName, Version)

	return bc
}
