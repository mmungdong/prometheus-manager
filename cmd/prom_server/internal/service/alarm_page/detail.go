package alarmPage

import (
	"context"
)

type (
	// DetailReq ...
	DetailReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
	}

	// DetailResp ...
	DetailResp struct {
		// add response params
	}
)

// Detail ...
func (l *AlarmPage) Detail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	// add your code here
	return &DetailResp{}, nil
}
