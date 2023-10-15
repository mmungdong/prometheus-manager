package main

import (
	"flag"

	ginplus "github.com/aide-cloud/gin-plus"
	"prometheus-manager/cmd/prom-server/internal/server"
	"prometheus-manager/pkg/hello"
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

	svs := []ginplus.Server{
		server.NewHttpServer(bc.Server),
	}

	hello.FmtASCIIGenerator(ServiceName, Version, bc.GetServer().GetMatadata())
	// 启动gin-plus
	ginplus.NewCtrlC(svs...).Start()
}
