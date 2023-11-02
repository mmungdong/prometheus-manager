package plog

import (
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

var _ log.Logger = (*Log)(nil)

type Log struct {
	logger *zap.Logger
}

// Log implements log.Logger.
func (*Log) Log(level log.Level, keyVales ...interface{}) error {
	switch level {
	case log.LevelDebug:
		ginplus.Logger().Sugar().Debug(keyVales...)
	case log.LevelInfo:
		ginplus.Logger().Sugar().Info(keyVales...)
	case log.LevelWarn:
		ginplus.Logger().Sugar().Warn(keyVales...)
	case log.LevelError:
		ginplus.Logger().Sugar().Error(keyVales...)
	case log.LevelFatal:
		ginplus.Logger().Sugar().Fatal(keyVales...)
	}
	return nil
}

func NewLog() *Log {
	return &Log{
		logger: ginplus.Logger(),
	}
}
