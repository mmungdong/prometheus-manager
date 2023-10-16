package conf

import (
	"sync"
	"time"

	"github.com/spf13/viper"
)

type (
	Bootstrap struct {
		// flag ...
		flag string

		// config ...
		// 存储配置
		Data *Data `yaml:"data"`
		// 服务配置
		Server *Server `yaml:"server"`
		// kafka配置
		Kafka *Kafka `yaml:"kafka"`
		// PushStrategy ...
		PushStrategy *PushStrategy `yaml:"push_strategy"`
		// Alert ...
		Alert *Alert `yaml:"alert"`
	}

	PushStrategy struct {
		Enable   bool          `yaml:"enable"`
		Intervel time.Duration `yaml:"intervel"`
		Topic    string        `yaml:"topic"`
	}

	Alert struct {
		Enable bool   `yaml:"enable"`
		Topic  string `yaml:"alert_topic"`
	}

	Data struct {
		// mysql ...
		Mysql *Mysql `yaml:"mysql"`
	}

	Server struct {
		Name     string            `yaml:"name"`
		Mode     string            `yaml:"mode"`
		Matadata map[string]string `yaml:"matadata"`

		// http ...
		Http *Http `yaml:"http"`
	}

	Http struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	Mysql struct {
		DSN   string `yaml:"dsn"`
		Debug bool   `yaml:"debug"`
	}

	Kafka struct {
		Enable    bool     `yaml:"enable"`
		Endpoints []string `yaml:"endpoints"`
	}

	BootstrapOption func(*Bootstrap)
)

var (
	c    *Bootstrap
	once sync.Once
)

func NewBootstrap(opts ...BootstrapOption) *Bootstrap {
	if c != nil {
		return c
	}

	b := &Bootstrap{
		flag: "config/config.yaml",
	}

	for _, o := range opts {
		o(b)
	}

	viper.SetConfigFile(b.flag)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(b); err != nil {
		panic(err)
	}

	// 写入完成, 清空viper
	viper.Reset()

	once.Do(func() {
		c = b
	})

	return b
}

func WithConfigFile(flag string) BootstrapOption {
	return func(b *Bootstrap) {
		b.flag = flag
	}
}
