package strategy

import (
	"context"
)

type (
	// DetailStrategyReq ...
	DetailStrategyReq struct {
		ID uint `uri:"id" binding:"required"`
	}

	// DetailStrategyResp ...
	DetailStrategyResp struct {
	}
)

func (l *Strategy) Detail(ctx context.Context, req *DetailStrategyReq) (*DetailStrategyResp, error) {
	// your code ...
	return &DetailStrategyResp{}, nil
}
