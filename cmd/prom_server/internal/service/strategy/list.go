package strategy

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	"prometheus-manager/pkg/model"
)

type (
	// ListStrategyReq ...
	ListStrategyReq struct {
		Keyword string `form:"keyword"`
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`
	}

	// ListStrategyResp ...
	ListStrategyResp struct {
		List []*model.PromStrategy `json:"list"`
		Page query.Page            `json:"page"`
	}
)

// List ...
func (l *Strategy) List(ctx context.Context, req *ListStrategyReq) (*ListStrategyResp, error) {
	return &ListStrategyResp{}, nil
}
