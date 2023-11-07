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
)

// NewDO 实例化prom_dict DO
func NewDO(values ...*model.PromDict) *object.DO[model.PromDict, Item] {
	return object.NewDO(
		object.DOWithList[model.PromDict, Item](values...),
		object.DOWithBuildFunc[model.PromDict, Item](modelToPromDict),
	)
}

// NewPO 实例化prom_dict PO
func NewPO(values ...*Item) *object.PO[model.PromDict, Item] {
	return object.NewPO(
		object.POWithList[model.PromDict, Item](values...),
		object.POWithBuildFunc[model.PromDict, Item](promDictToModel),
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
