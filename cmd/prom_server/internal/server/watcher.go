package server

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-kratos/kratos/v2/log"

	"prometheus-manager/cmd/prom_server/internal/conf"
	"prometheus-manager/cmd/prom_server/internal/service/alert"
	"prometheus-manager/pkg/conn"
)

type WatchServer struct {
	watchService *alert.Alert
	logger       *log.Helper
	consumer     *conn.KafkaConsumer

	kafkaConf *conf.Kafka
	alertConf *conf.Alert
}

var _ ginplus.Server = (*WatchServer)(nil)

func NewWatchServer(bc *conf.Bootstrap, logger log.Logger) *WatchServer {
	kafkaConf := bc.GetKafka()
	alertConf := bc.GetAlert()
	watchService := alert.NewAlert()
	logHelper := log.NewHelper(log.With(logger, "module", "server/server"))
	if !kafkaConf.GetEnable() {
		logHelper.Warnf("Not enable kafka")
		return &WatchServer{
			logger:    logHelper,
			alertConf: alertConf,
			kafkaConf: kafkaConf,
		}
	}

	consumer, err := conn.NewKafkaConsumer(kafkaConf.GetEndpoints(), []string{alertConf.GetAlertTopic()}, log.DefaultLogger)
	if err != nil {
		logHelper.Error("kafka消费者初始化失败")
		panic(err)
	}

	return &WatchServer{
		watchService: watchService,
		logger:       logHelper,
		consumer:     consumer,
		alertConf:    alertConf,
		kafkaConf:    kafkaConf,
	}
}

func (l *WatchServer) Start() error {
	l.logger.Info("[WatchServer] server starting")

	// 如果没有启用监听，则直接返回
	if !l.alertConf.GetEnable() {
		return nil
	}
	if l.consumer != nil {
		l.consumer.Consume(func(msg *kafka.Message) bool {
			ctx, cancel := context.WithTimeoutCause(context.Background(), 5*time.Second, errors.New("kafka消费超时"))
			defer cancel()
			return l.callbackFunc(ctx, msg)
		})
	}

	return nil
}

func (l *WatchServer) Stop() {
	defer l.logger.Info("[WatchServer] server stopped")
	if l.consumer != nil {
		_ = l.consumer.Close()
	}
}

func (l *WatchServer) callbackFunc(ctx context.Context, msg *kafka.Message) (flag bool) {
	flag = true
	l.logger.Infof("message at topic:%v time:%d %s = %s\n", msg.TopicPartition, msg.Timestamp.Unix(), string(msg.Key), string(msg.Value))

	var req alert.HookReq
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		l.logger.Errorf("kafka消息错误: %v", err)
		return
	}

	hookResp, err := l.watchService.PostHook(ctx, &req)
	if err != nil {
		l.logger.Errorf("kafka消费失败: %v", err)
		return
	}
	l.logger.Infof("kafka消费成功, %v", hookResp)
	return
}
