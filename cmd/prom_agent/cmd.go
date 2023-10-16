package main

import (
	"flag"

	"prometheus-manager/cmd/prom_agent/internal/server"
	"prometheus-manager/pkg/hello"
	"prometheus-manager/pkg/plog"

	ginplus "github.com/aide-cloud/gin-plus"
)

var (
	// Version 版本号
	Version string
	// ServiceName 服务名称
	ServiceName = "prom-server"
	// configPath 配置文件路径
	configPath = flag.String("config", "configs/config.yaml", "config file path")
)

func main() {
	flag.Parse()
	bc := Init()
	// 初始化日志
	logger := plog.NewLog()

	hello.FmtASCIIGenerator(ServiceName, Version, bc.GetServer().GetMatadata())

	svs := []ginplus.Server{
		server.NewHttpServer(bc.GetServer(), logger),
	}

	// 启动gin-plus
	ginplus.NewCtrlC(svs...).Start()
}
