package strategyGroup

import (
	"context"

	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
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

	categories := make([]*model.PromDict, 0, len(req.CategoriesIds))
	for _, categoryId := range req.CategoriesIds {
		categories = append(categories, &model.PromDict{BaseModel: query.BaseModel{ID: categoryId}})
	}

	newStrategyGroup := &model.PromGroup{
		Name:   req.Name,
		Remark: req.Remark,

		Categories: categories,
	}

	if err := strategyGroupData.WithContext(ctx).Create(newStrategyGroup); err != nil {
		return nil, err
	}
	// add your code here
	return &CreateResp{ID: newStrategyGroup.ID}, nil
}
