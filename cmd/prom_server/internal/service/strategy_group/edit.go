package strategyGroup

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
	"prometheus-manager/pkg/model"
)

type (
	// EditReq ...
	EditReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`

		Name   string `json:"name"`
		Remark string `json:"remark"`

		Status        model.Status `json:"status"`
		CategoriesIds []uint       `json:"categories_ids"`
	}

	// EditResp ...
	EditResp struct {
		// add response params
		ID uint `json:"id"`
	}
)

// Edit ...
func (l *StrategyGroup) Edit(ctx context.Context, req *EditReq) (*EditResp, error) {
	strategyGroupData := dataStrategyGroup.NewStrategyGroup()

	// 查询策略组
	strategyGroup, err := strategyGroupData.WithContext(ctx).FirstByID(req.ID)
	if err != nil {
		return nil, err
	}

	categories := make([]schema.Tabler, 0, len(req.CategoriesIds))
	for _, categoryId := range req.CategoriesIds {
		categories = append(categories, &model.PromDict{BaseModel: query.BaseModel{ID: categoryId}})
	}

	updateMapData := buildEditStrategyGroupMap(req, strategyGroup)

	err = strategyGroupData.DB().Transaction(func(tx *gorm.DB) error {
		if len(updateMapData) > 0 {
			if err = strategyGroupData.WithDB(tx).WithContext(ctx).UpdateMapByID(req.ID, updateMapData); err != nil {
				return err
			}
		}

		// 替换关联
		if err = strategyGroupData.WithDB(tx).WithContext(ctx).AssociationReplace(query.AssociationKey(strategyGroupData.PreloadCategoriesKey), categories...); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &EditResp{ID: req.ID}, nil
}

// buildEditStrategyGroupMap ...
func buildEditStrategyGroupMap(req *EditReq, strategyGroup *model.PromGroup) map[string]interface{} {
	updateMapData := make(map[string]interface{})

	if req.Name != strategyGroup.Name {
		updateMapData["name"] = req.Name
	}

	if req.Remark != strategyGroup.Remark {
		updateMapData["remark"] = req.Remark
	}

	if req.Status != strategyGroup.Status {
		updateMapData["status"] = req.Status
	}

	return updateMapData
}
