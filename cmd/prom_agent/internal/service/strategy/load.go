package strategy

import (
	"context"
)

type (
	// LoadReq ...
	LoadReq struct {
		// add request params
	}

	// LoadResp ...
	LoadResp struct {
		// add response params
	}
)

// PostLoad ...
func (l *Strategy) PostLoad(ctx context.Context, req *LoadReq) (*LoadResp, error) {
	// add your code here
	return &LoadResp{}, nil
}
