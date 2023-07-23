package conf

import (
	"github.com/google/wire"
	"prometheus-manager/pkg/conn"
	"sync"
)

// ProviderSet is conf providers.
var ProviderSet = wire.NewSet(
	wire.FieldsOf(new(*Bootstrap), "Trace"),
	wire.FieldsOf(new(*Bootstrap), "Server"),
	wire.FieldsOf(new(*Bootstrap), "Strategy"),
	wire.FieldsOf(new(*Bootstrap), "Env"),
	wire.Bind(new(conn.ITraceConfig), new(*Trace)),
	wire.Bind(new(conn.ITraceEnv), new(*Env)),
)

var _ conn.ITraceConfig = (*Trace)(nil)
var _ conn.ITraceEnv = (*Env)(nil)

// 全局配置获取
var c *Bootstrap

// 保证只执行一次
var once sync.Once

// Get 获取配置
func Get() *Bootstrap {
	return c
}

// Set 设置配置, 单例
func Set(b *Bootstrap) {
	once.Do(func() {
		c = b
	})
}