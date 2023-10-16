package server

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"prometheus-manager/cmd/prom_agent/internal/conf"
	"prometheus-manager/cmd/prom_agent/internal/service/strategy"
	"prometheus-manager/pkg/conn"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-kratos/kratos/v2/log"
)

var _ ginplus.Server = (*WatchServer)(nil)

type WatchServer struct {
	watchService *strategy.Strategy
	logger       *log.Helper
	consumer     *conn.KafkaConsumer

	kafkaConf         *conf.Kafka
	synceStrategyConf *conf.SynceStrategy
}

func NewWatchServer(kafkaConf *conf.Kafka, synceStrategyConf *conf.SynceStrategy, logger log.Logger) *WatchServer {
	watchService := strategy.NewStrategy()
	logHelper := log.NewHelper(log.With(logger, "module", "server/server"))
	if !kafkaConf.GetEnable() {
		logHelper.Warnf("Not enabel kafka")
		return &WatchServer{
			logger:            logHelper,
			synceStrategyConf: synceStrategyConf,
			kafkaConf:         kafkaConf,
		}
	}

	consumer, err := conn.NewKafkaConsumer(kafkaConf.GetEndpoints(), []string{synceStrategyConf.GetTopic()}, log.DefaultLogger)
	if err != nil {
		logHelper.Error("kafka消费者初始化失败")
		panic(err)
	}

	return &WatchServer{
		watchService:      watchService,
		logger:            logHelper,
		consumer:          consumer,
		synceStrategyConf: synceStrategyConf,
		kafkaConf:         kafkaConf,
	}
}

// Start implements ginplus.Server.
func (l *WatchServer) Start() error {
	l.logger.Info("[WatchServer] server starting")

	// 如果没有启用监听，则直接返回
	if !l.synceStrategyConf.GetEnable() {
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
		l.consumer.Close()
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
