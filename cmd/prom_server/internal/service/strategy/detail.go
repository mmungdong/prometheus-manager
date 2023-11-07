package strategy

import (
	"context"

	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/pkg/model"
)

type (
	// DetailStrategyReq ...
	DetailStrategyReq struct {
		ID uint `uri:"id" binding:"required"`

		IsDelete bool `form:"is_delete"`
	}

	// DetailStrategyResp ...
	DetailStrategyResp struct {
		*Item
	}
)

func (l *Strategy) Detail(ctx context.Context, req *DetailStrategyReq) (*DetailStrategyResp, error) {
	strategyData := dataStrategy.NewStrategy()

	var detailInfo *model.PromStrategy
	var err error
	if req.IsDelete {
		detailInfo, err = strategyData.WithContext(ctx).FirstByIDWithTrashed(req.ID)
	} else {
		detailInfo, err = strategyData.WithContext(ctx).FirstByID(req.ID)
	}

	if err != nil {
		return nil, err
	}

	return &DetailStrategyResp{
		Item: NewDO(detailInfo).PO().One(),
	}, nil
}
