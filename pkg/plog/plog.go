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
func (*Log) Log(level log.Level, keyvals ...interface{}) error {
	switch level {
	case log.LevelDebug:
		ginplus.Logger().Sugar().Debug(keyvals...)
	case log.LevelInfo:
		ginplus.Logger().Sugar().Info(keyvals...)
	case log.LevelWarn:
		ginplus.Logger().Sugar().Warn(keyvals...)
	case log.LevelError:
		ginplus.Logger().Sugar().Error(keyvals...)
	case log.LevelFatal:
		ginplus.Logger().Sugar().Fatal(keyvals...)
	}
	return nil
}

func NewLog() *Log {
	return &Log{
		logger: ginplus.Logger(),
	}
}
