package alarmHistory

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
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
		Page query.Page `json:"page"`
	}
)

// List ...
func (l *AlarmHistory) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	// add your code here
	return &ListResp{}, nil
}
