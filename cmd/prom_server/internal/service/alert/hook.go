package alert

import (
	"context"
)

type (
	// HookReq ...
	HookReq struct {
		// add request params
	}

	// HookResp ...
	HookResp struct {
		// add response params
	}
)

// PostHook ...
func (l *Alert) PostHook(ctx context.Context, req *HookReq) (*HookResp, error) {
	// add your code here
	return &HookResp{}, nil
}
