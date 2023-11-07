package strategyGroup

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/plugin/soft_delete"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/object"
	"prometheus-manager/pkg/times"
)

type (
	// Item ...
	Item struct {
		ID            uint                  `json:"id"`
		Name          string                `json:"name"`
		Remark        string                `json:"remark"`
		StrategyCount int64                 `json:"strategy_count"`
		Status        model.Status          `json:"status"`
		CreatedAt     int64                 `json:"created_at"`
		UpdatedAt     int64                 `json:"updated_at"`
		DeletedAt     soft_delete.DeletedAt `json:"deleted_at"`

		CategoriesIds []uint      `json:"categories_ids"`
		Categories    []*DictItem `json:"categories"`
	}

	DictItem struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}
)

// NewDO 实例化strategyGroup DO
func NewDO(values ...*model.PromGroup) *object.DO[model.PromGroup, Item] {
	return object.NewDO(
		object.DOWithList[model.PromGroup, Item](values...),
		object.DOWithBuildFunc[model.PromGroup, Item](modelToStrategyGroup),
	)
}

// NewPO 实例化strategyGroup PO
func NewPO(values ...*Item) *object.PO[model.PromGroup, Item] {
	return object.NewPO(
		object.POWithList[model.PromGroup, Item](values...),
		object.POWithBuildFunc[model.PromGroup, Item](strategyGroupToModel),
	)
}

// modelToStrategyGroup ...
func modelToStrategyGroup(model *model.PromGroup) *Item {
	if model == nil {
		return nil
	}

	categories := make([]*DictItem, 0, len(model.Categories))
	categoriesIds := make([]uint, 0, len(model.Categories))
	for _, category := range model.Categories {
		categories = append(categories, &DictItem{
			ID:    category.ID,
			Name:  category.Name,
			Color: category.Color,
		})
		categoriesIds = append(categoriesIds, category.ID)
	}

	return &Item{
		ID:            model.ID,
		Name:          model.Name,
		Remark:        model.Remark,
		StrategyCount: model.StrategyCount,
		Status:        model.Status,
		CreatedAt:     times.TimeToUnix(model.CreatedAt),
		UpdatedAt:     times.TimeToUnix(model.UpdatedAt),
		DeletedAt:     model.DeletedAt,
		Categories:    categories,
		CategoriesIds: categoriesIds,
	}
}

// strategyGroupToModel ...
func strategyGroupToModel(item *Item) *model.PromGroup {
	if item == nil {
		return nil
	}

	categories := make([]*model.PromDict, 0, len(item.CategoriesIds))
	for _, categoryId := range item.CategoriesIds {
		categories = append(categories, &model.PromDict{BaseModel: query.BaseModel{ID: categoryId}})
	}

	return &model.PromGroup{
		BaseModel: query.BaseModel{
			ID:        item.ID,
			CreatedAt: times.UnixToTime(item.CreatedAt),
			UpdatedAt: times.UnixToTime(item.UpdatedAt),
			DeletedAt: item.DeletedAt,
		},
		Name:          item.Name,
		Remark:        item.Remark,
		StrategyCount: item.StrategyCount,
		Status:        item.Status,
		Categories:    categories,
	}
}
