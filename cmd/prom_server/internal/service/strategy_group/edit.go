package strategyGroup

import (
	"context"
)

type (
	// EditReq ...
	EditReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
	}

	// EditResp ...
	EditResp struct {
		// add response params
	}
)

// Edit ...
func (l *StrategyGroup) Edit(ctx context.Context, req *EditReq) (*EditResp, error) {
	// add your code here
	return &EditResp{}, nil
}
