package strategyGroup

import (
	"context"

	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
)

type (
	// DeleteReq ...
	DeleteReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
	}

	// DeleteResp ...
	DeleteResp struct {
		ID uint `json:"id"`
	}
)

// Delete ...
func (l *StrategyGroup) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	strategyGroupData := dataStrategyGroup.NewStrategyGroup()

	if err := strategyGroupData.WithContext(ctx).DeleteByID(req.ID); err != nil {
		return nil, err
	}
	// add your code here
	return &DeleteResp{ID: req.ID}, nil
}
