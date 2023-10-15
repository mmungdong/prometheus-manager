package strategy

import (
	"context"
)

type (
	// DeleteStrategyReq ...
	DeleteStrategyReq struct {
		ID uint `uri:"id" binding:"required"`
	}

	// DeleteStrategyResp ...
	DeleteStrategyResp struct {
	}
)

func (l *Strategy) Delete(ctx context.Context, req *DeleteStrategyReq) (*DeleteStrategyResp, error) {
	// your code ...
	return &DeleteStrategyResp{}, nil
}
