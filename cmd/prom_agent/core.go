package main

import (
	ginplus "github.com/aide-cloud/gin-plus"
	
	"prometheus-manager/cmd/prom_agent/internal/conf"
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

	return bc
}
