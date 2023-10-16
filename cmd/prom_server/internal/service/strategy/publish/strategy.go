package publish

import (
	"context"
	"time"
)

type (
	// StrategyReq ...
	StrategyReq struct {
		// add request params
		Topic string `json:"topic"`
	}

	// StrategyResp ...
	StrategyResp struct {
		// add response params
		Msg       string `json:"msg"`
		Timestamp int64  `json:"timestamp"`
	}
)

// PostStrategy ...
func (l *Publish) PostStrategy(ctx context.Context, req *StrategyReq) (*StrategyResp, error) {
	// 发布策略通知
	// add your code here
	return &StrategyResp{
		Msg:       "success",
		Timestamp: time.Now().Unix(),
	}, nil
}
