package strategy

import (
	"context"
)

type (
	// CreateStrategyReq ...
	CreateStrategyReq struct {
	}

	// CreateStrategyResp ...
	CreateStrategyResp struct {
	}
)

func (l *Strategy) Create(ctx context.Context, req *CreateStrategyReq) (*CreateStrategyResp, error) {
	// your code ...
	return &CreateStrategyResp{}, nil
}
