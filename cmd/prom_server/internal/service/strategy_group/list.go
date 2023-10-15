package strategyGroup

import (
	"context"
)

type (
	// ListReq ...
	ListReq struct {
		// add request params
	}

	// ListResp ...
	ListResp struct {
		// add response params
	}
)

// PostList ...
func (l *StrategyGroup) PostList(ctx context.Context, req *ListReq) (*ListResp, error) {
	// add your code here
	return &ListResp{}, nil
}
