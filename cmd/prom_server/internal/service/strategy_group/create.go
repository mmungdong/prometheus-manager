package strategyGroup

import (
	"context"

	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
)

type (
	// CreateReq ...
	CreateReq struct {
		Name   string `json:"name"`
		Remark string `json:"remark"`

		CategoriesIds []uint `json:"categories_ids"`
	}

	// CreateResp ...
	CreateResp struct {
		ID uint `json:"id"`
	}
)

// Create ...
func (l *StrategyGroup) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	strategyGroupData := dataStrategyGroup.NewStrategyGroup()

	newStrategyGroupPOValue := &Item{
		Name:   req.Name,
		Remark: req.Remark,

		CategoriesIds: req.CategoriesIds,
	}
	newStrategyGroup := NewPO(newStrategyGroupPOValue).DO().One()

	if err := strategyGroupData.WithContext(ctx).Create(newStrategyGroup); err != nil {
		return nil, err
	}
	// add your code here
	return &CreateResp{ID: newStrategyGroup.ID}, nil
}
