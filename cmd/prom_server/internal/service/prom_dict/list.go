package promDict

import (
	"context"
)

type (
	// ListReq ...
	ListReq struct {
		// add request params
		Keyword string `form:"keyword"`
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`
	}

	// ListResp ...
	ListResp struct {
		// add response params
	}
)

// List ...
func (l *PromDict) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	// add your code here
	return &ListResp{}, nil
}
