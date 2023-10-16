package publish

import (
	"context"
)

type (
	// StrategyReq ...
	StrategyReq struct {
		// add request params
	}

	// StrategyResp ...
	StrategyResp struct {
		// add response params
	}
)

// PostStrategy ...
func (l *Publish) PostStrategy(ctx context.Context, req *StrategyReq) (*StrategyResp, error) {
	// 发布策略通知
	// add your code here
	return &StrategyResp{}, nil
}
