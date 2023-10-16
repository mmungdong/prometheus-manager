package alert

import (
	"context"
	"prometheus-manager/pkg/alert"
)

type (
	// HookReq ...
	HookReq struct {
		// add request params
		*alert.Data
	}

	// HookResp ...
	HookResp struct {
		// add response params
		Msg string `json:"msg"`
	}
)

// PostHook ...
func (l *Alert) PostHook(ctx context.Context, req *HookReq) (*HookResp, error) {
	// add your code here
	return &HookResp{}, nil
}
