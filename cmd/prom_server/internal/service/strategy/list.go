package strategy

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
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
		List []*Item     `json:"list"`
		Page *query.Page `json:"page"`
	}
)

// List ...
func (l *Strategy) List(ctx context.Context, req *ListStrategyReq) (*ListStrategyResp, error) {
	strategyData := dataStrategy.NewStrategy()

	wheres := []query.ScopeMethod{
		query.WhereLikeKeyword(req.Keyword, "alert"),
		strategyData.PreloadAlertLevel(),
		strategyData.PreloadGroupInfo(),
	}

	pgInfo := query.NewPage(req.Curr, req.Size)

	strategies, err := strategyData.WithContext(ctx).List(pgInfo, wheres...)
	if err != nil {
		return nil, err
	}

	list := NewDO(strategies...).PO().List()

	return &ListStrategyResp{
		List: list,
		Page: pgInfo,
	}, nil
}
