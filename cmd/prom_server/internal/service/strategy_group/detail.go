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
		ID             uint                `json:"id"`
		Name           string              `json:"name"`
		Remark         string              `json:"remark"`
		StrategyCount  int64               `json:"strategy_count"`
		Status         model.Status        `json:"status"`
		PromStrategies []*PromStrategyItem `json:"prom_strategies"`
		Categories     []*CategoryItem     `json:"categories"`
	}

	PromStrategyItem struct {
		ID    uint   `json:"id"`
		Alert string `json:"alert"`
	}

	CategoryItem struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
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
		ID:             strategyGroup.ID,
		Name:           strategyGroup.Name,
		Remark:         strategyGroup.Remark,
		StrategyCount:  strategyGroup.StrategyCount,
		Status:         strategyGroup.Status,
		PromStrategies: buildPromStrategies(strategyGroup.PromStrategies),
		Categories:     buildCategories(strategyGroup.Categories),
	}, nil
}

func buildPromStrategies(promStrategies []*model.PromStrategy) []*PromStrategyItem {
	items := make([]*PromStrategyItem, 0, len(promStrategies))
	for _, promStrategy := range promStrategies {
		items = append(items, &PromStrategyItem{
			ID:    promStrategy.ID,
			Alert: promStrategy.Alert,
		})
	}
	return items
}

func buildCategories(categories []*model.PromDict) []*CategoryItem {
	items := make([]*CategoryItem, 0, len(categories))
	for _, category := range categories {
		items = append(items, &CategoryItem{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return items
}
