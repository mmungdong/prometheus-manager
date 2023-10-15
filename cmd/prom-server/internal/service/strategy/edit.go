package strategy

import (
	"context"
)

type (
	// EditStrategyReq ...
	EditStrategyReq struct {
		ID uint `uri:"id" binding:"required"`
	}

	// EditStrategyResp ...
	EditStrategyResp struct {
	}
)

func (l *Strategy) Edit(ctx context.Context, req *EditStrategyReq) (*EditStrategyResp, error) {
	// your code ...
	return &EditStrategyResp{}, nil
}
