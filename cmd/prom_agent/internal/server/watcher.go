package server

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-kratos/kratos/v2/log"

	"prometheus-manager/cmd/prom_agent/internal/conf"
	"prometheus-manager/cmd/prom_agent/internal/service/strategy"
	"prometheus-manager/pkg/conn"
)

var _ ginplus.Server = (*WatchServer)(nil)

type WatchServer struct {
	watchService *strategy.Strategy
	logger       *log.Helper
	consumer     *conn.KafkaConsumer

	kafkaConf        *conf.Kafka
	syncStrategyConf *conf.SyncStrategy
}

// NewWatchServer 初始化监听mq消息的服务(kafka)
func NewWatchServer(bc *conf.Bootstrap, logger log.Logger) *WatchServer {
	kafkaConf := bc.GetKafka()
	syncStrategyConf := bc.GetSyncStrategy()
	watchService := strategy.NewStrategy()
	logHelper := log.NewHelper(log.With(logger, "module", "server/server"))
	if !kafkaConf.GetEnable() {
		logHelper.Warnf("Not enable kafka")
		return &WatchServer{
			logger:           logHelper,
			syncStrategyConf: syncStrategyConf,
			kafkaConf:        kafkaConf,
		}
	}

	consumer, err := conn.NewKafkaConsumer(kafkaConf.GetEndpoints(), []string{syncStrategyConf.GetTopic()}, log.DefaultLogger)
	if err != nil {
		logHelper.Error("kafka消费者初始化失败")
		panic(err)
	}

	return &WatchServer{
		watchService:     watchService,
		logger:           logHelper,
		consumer:         consumer,
		syncStrategyConf: syncStrategyConf,
		kafkaConf:        kafkaConf,
	}
}

// Start implements ginplus.Server.
func (l *WatchServer) Start() error {
	l.logger.Info("[WatchServer] server starting")

	// 如果没有启用监听，则直接返回
	if !l.syncStrategyConf.GetEnable() {
		return nil
	}

	// 启动kafka消费者
	if l.consumer != nil {
		l.consumer.Consume(func(msg *kafka.Message) (flag bool) {
			ctx, cancel := context.WithTimeoutCause(context.Background(), 5*time.Second, errors.New("kafka消费超时"))
			defer cancel()
			flag = l.callbackFunc(ctx, msg)
			return
		})
	}

	return nil
}

// Stop implements ginplus.Server.
func (l *WatchServer) Stop() {
	defer l.logger.Info("[WatchServer] server stopped")
	if l.consumer != nil {
		_ = l.consumer.Close()
	}
}

func (l *WatchServer) callbackFunc(ctx context.Context, msg *kafka.Message) (flag bool) {
	flag = true
	l.logger.Infof("message at topic:%v time:%d %s = %s\n", msg.TopicPartition, msg.Timestamp.Unix(), string(msg.Key), string(msg.Value))

	var req strategy.LoadReq
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		l.logger.Errorf("kafka消息错误: %v", err)
		return
	}

	alert, err := l.watchService.PostLoad(ctx, &req)
	if err != nil {
		l.logger.Errorf("kafka消费失败: %v", err)
		return
	}
	l.logger.Infof("kafka消费成功, %v", alert)
	return
}
