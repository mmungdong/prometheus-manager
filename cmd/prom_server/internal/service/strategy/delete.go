package strategy

import (
	"context"

	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
)

type (
	// DeleteStrategyReq ...
	DeleteStrategyReq struct {
		ID uint `uri:"id" binding:"required"`
	}

	// DeleteStrategyResp ...
	DeleteStrategyResp struct {
		ID uint `json:"id"`
	}
)

func (l *Strategy) Delete(ctx context.Context, req *DeleteStrategyReq) (*DeleteStrategyResp, error) {
	strategyData := dataStrategy.NewStrategy()

	if err := strategyData.WithContext(ctx).DeleteByID(req.ID); err != nil {
		return nil, err
	}
	return &DeleteStrategyResp{ID: req.ID}, nil
}
