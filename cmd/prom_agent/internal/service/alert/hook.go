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
	}
)

// PostHook ...
func (l *Alert) PostHook(ctx context.Context, req *HookReq) (*HookResp, error) {
	// 通过kafka发送alert消息
	// add your code here
	return &HookResp{}, nil
}
