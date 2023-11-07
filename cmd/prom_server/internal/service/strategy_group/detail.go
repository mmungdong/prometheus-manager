package strategyGroup

import (
	"context"

	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
	"prometheus-manager/pkg/model"
)

type (
	// DetailReq ...
	DetailReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`

		IsDelete bool `form:"is_delete"`
	}

	// DetailResp ...
	DetailResp struct {
		*Item
	}
)

// Detail ...
func (l *StrategyGroup) Detail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	strategyGroupData := dataStrategyGroup.NewStrategyGroup()

	var strategyGroup *model.PromGroup
	var err error

	if req.IsDelete {
		strategyGroup, err = strategyGroupData.WithContext(ctx).FirstByIDWithTrashed(req.ID)
	} else {
		strategyGroup, err = strategyGroupData.WithContext(ctx).FirstByID(req.ID)
	}

	if err != nil {
		return nil, err
	}
	return &DetailResp{
		Item: NewDO(strategyGroup).PO().One(),
	}, nil
}
