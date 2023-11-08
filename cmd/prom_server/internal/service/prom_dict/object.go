package promDict

import (
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/object"
	"prometheus-manager/pkg/times"
)

type (
	Item struct {
		ID        uint  `json:"id"`
		CreatedAt int64 `json:"create_at"`
		UpdatedAt int64 `json:"update_at"`
		DeletedAt int64 `json:"delete_at"`

		Name     string         `json:"name"`     // 字典名称
		Category model.Category `json:"category"` // 字典类型
		Color    string         `json:"color"`    // 字典tag颜色
		Status   model.Status   `json:"status"`   // 状态
		Remark   string         `json:"remark"`   // 字典备注
	}

	// SelectItem ...
	SelectItem struct {
		Label    string         `json:"label"`    // 对应前端的label, 为Name
		Value    any            `json:"value"`    // 对应前端的value, 为ID
		Category model.Category `json:"category"` // 字典类型
		Color    string         `json:"color"`    // 字典tag颜色
		Status   model.Status   `json:"status"`   // 状态
		Remark   string         `json:"remark"`   // 字典备注
	}
)

// NewDO 实例化prom_dict DO
func NewDO(values ...*model.PromDict) *object.DO[model.PromDict, Item] {
	return object.NewDO(
		object.DOWithList[model.PromDict, Item](values...),
		object.DOWithBuildFunc[model.PromDict, Item](modelToPromDict),
	)
}

// NewSelectDO 实例化prom_dict DO
func NewSelectDO(values ...*model.PromDict) *object.DO[model.PromDict, SelectItem] {
	return object.NewDO(
		object.DOWithList[model.PromDict, SelectItem](values...),
		object.DOWithBuildFunc[model.PromDict, SelectItem](modelToSelectItem),
	)
}

// NewPO 实例化prom_dict PO
func NewPO(values ...*Item) *object.PO[model.PromDict, Item] {
	return object.NewPO(
		object.POWithList[model.PromDict, Item](values...),
		object.POWithBuildFunc[model.PromDict, Item](promDictToModel),
	)
}

// NewSelectPO 实例化prom_dict PO
func NewSelectPO(values ...*SelectItem) *object.PO[model.PromDict, SelectItem] {
	return object.NewPO(
		object.POWithList[model.PromDict, SelectItem](values...),
		object.POWithBuildFunc[model.PromDict, SelectItem](selectItemToModel),
	)
}

// promDictToModel ...
func promDictToModel(promDict *Item) *model.PromDict {
	return &model.PromDict{
		Name:     promDict.Name,
		Category: promDict.Category,
		Color:    promDict.Color,
		Status:   promDict.Status,
		Remark:   promDict.Remark,
	}
}

// selectItemToModel ...
func selectItemToModel(selectItem *SelectItem) *model.PromDict {
	return &model.PromDict{
		Name:     selectItem.Label,
		Category: selectItem.Category,
		Color:    selectItem.Color,
		Status:   selectItem.Status,
		Remark:   selectItem.Remark,
	}
}

// modelToPromDict ...
func modelToPromDict(modelItem *model.PromDict) *Item {
	return &Item{
		ID:        modelItem.ID,
		CreatedAt: times.TimeToUnix(modelItem.CreatedAt),
		UpdatedAt: times.TimeToUnix(modelItem.UpdatedAt),
		DeletedAt: int64(modelItem.DeletedAt),
		Name:      modelItem.Name,
		Category:  modelItem.Category,
		Color:     modelItem.Color,
		Status:    modelItem.Status,
		Remark:    modelItem.Remark,
	}
}

// modelToSelectItem ...
func modelToSelectItem(modelItem *model.PromDict) *SelectItem {
	return &SelectItem{
		Label:    modelItem.Name,
		Value:    modelItem.ID,
		Category: modelItem.Category,
		Color:    modelItem.Color,
		Status:   modelItem.Status,
		Remark:   modelItem.Remark,
	}
}
