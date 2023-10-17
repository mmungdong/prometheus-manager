package alert

import (
	"context"
	"prometheus-manager/pkg/alert"

	ginplus "github.com/aide-cloud/gin-plus"
	"go.uber.org/zap"
)

type (
	// HookReq ...
	HookReq struct {
		*alert.Data
	}
	// HookResp ...
	HookResp struct {
		*alert.Data
	}
)

// PostHook ...
func (l *Alert) PostHook(ctx context.Context, req *HookReq) (*HookResp, error) {
	ginplus.Logger().Info("PostHook: ", zap.Any("req", req))
	// 通过kafka发送alert消息
	// add your code here
	return &HookResp{
		Data: req.Data,
	}, nil
}
