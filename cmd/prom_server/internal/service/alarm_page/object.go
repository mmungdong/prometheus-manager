package alarmPage

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/plugin/soft_delete"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/object"
	"prometheus-manager/pkg/times"
)

type (
	Item struct {
		ID     uint         `json:"id"`
		Name   string       `json:"name"`
		Icon   string       `json:"icon"`
		Color  string       `json:"color"`
		Remark string       `json:"remark"`
		Status model.Status `json:"status"`

		CreatedAt      int64                 `json:"created_at"`
		UpdatedAt      int64                 `json:"updated_at"`
		DeletedAt      soft_delete.DeletedAt `json:"deleted_at"`
		PromStrategies []*PromStrategyItem   `json:"prom_strategies"`

		promStrategyIds []uint
	}

	PromStrategyItem struct {
		ID    uint   `json:"id"`
		Alert string `json:"alert"`
	}

	SelectItem struct {
		Label  string       `json:"label"` // 对应前端的label, 为Name
		Value  any          `json:"value"` // 对应前端的value, 为ID
		Icon   string       `json:"icon"`
		Color  string       `json:"color"`
		Remark string       `json:"remark"`
		Status model.Status `json:"status"`
	}
)

// NewDO 实例化alarmPage DO
func NewDO(values ...*model.PromAlarmPage) *object.DO[model.PromAlarmPage, Item] {
	return object.NewDO(
		object.DOWithList[model.PromAlarmPage, Item](values...),
		object.DOWithBuildFunc[model.PromAlarmPage, Item](modelToAlarmPage),
	)
}

// NewSelectDO 实例化alarmPage DO
func NewSelectDO(values ...*model.PromAlarmPage) *object.DO[model.PromAlarmPage, SelectItem] {
	return object.NewDO(
		object.DOWithList[model.PromAlarmPage, SelectItem](values...),
		object.DOWithBuildFunc[model.PromAlarmPage, SelectItem](modelToSelectItem),
	)
}

// NewPO 实例化alarmPage PO
func NewPO(values ...*Item) *object.PO[model.PromAlarmPage, Item] {
	return object.NewPO(
		object.POWithList[model.PromAlarmPage, Item](values...),
		object.POWithBuildFunc[model.PromAlarmPage, Item](alarmPageToModel),
	)
}

// NewSelectPO 实例化alarmPage PO
func NewSelectPO(values ...*SelectItem) *object.PO[model.PromAlarmPage, SelectItem] {
	return object.NewPO(
		object.POWithList[model.PromAlarmPage, SelectItem](values...),
		object.POWithBuildFunc[model.PromAlarmPage, SelectItem](selectItemToModel),
	)
}

// alarmPageToModel ...
func alarmPageToModel(item *Item) *model.PromAlarmPage {

	return &model.PromAlarmPage{
		Name:   item.Name,
		Icon:   item.Icon,
		Color:  item.Color,
		Remark: item.Remark,
		Status: item.Status,

		PromStrategies: buildModelPromStrategies(item.promStrategyIds),

		BaseModel: query.BaseModel{
			ID:        item.ID,
			CreatedAt: times.UnixToTime(item.CreatedAt),
			UpdatedAt: times.UnixToTime(item.UpdatedAt),
			DeletedAt: item.DeletedAt,
		},
	}
}

// selectItemToModel ...
func selectItemToModel(selectItem *SelectItem) *model.PromAlarmPage {
	return &model.PromAlarmPage{
		Name:   selectItem.Label,
		Color:  selectItem.Color,
		Remark: selectItem.Remark,
		Status: selectItem.Status,
		Icon:   selectItem.Icon,
		BaseModel: query.BaseModel{
			ID: selectItem.Value.(uint),
		},
	}
}

// modelToAlarmPage ...
func modelToAlarmPage(alarmPage *model.PromAlarmPage) *Item {
	return &Item{
		ID:     alarmPage.ID,
		Name:   alarmPage.Name,
		Icon:   alarmPage.Icon,
		Color:  alarmPage.Color,
		Remark: alarmPage.Remark,
		Status: alarmPage.Status,

		PromStrategies: buildPromStrategies(alarmPage.PromStrategies),

		CreatedAt: times.TimeToUnix(alarmPage.CreatedAt),
		UpdatedAt: times.TimeToUnix(alarmPage.UpdatedAt),
		DeletedAt: alarmPage.DeletedAt,
	}
}

// modelToSelectItem ...
func modelToSelectItem(alarmPage *model.PromAlarmPage) *SelectItem {
	return &SelectItem{
		Label:  alarmPage.Name,
		Value:  alarmPage.ID,
		Icon:   alarmPage.Icon,
		Color:  alarmPage.Color,
		Remark: alarmPage.Remark,
		Status: alarmPage.Status,
	}
}

// buildModelPromStrategies ...
func buildModelPromStrategies(promStrategies []uint) []*model.PromStrategy {
	list := make([]*model.PromStrategy, 0, len(promStrategies))
	for _, promStrategyId := range promStrategies {
		list = append(list, &model.PromStrategy{BaseModel: query.BaseModel{ID: promStrategyId}})
	}
	return list
}

// buildPromStrategies ...
func buildPromStrategies(promStrategies []*model.PromStrategy) []*PromStrategyItem {
	list := make([]*PromStrategyItem, 0, len(promStrategies))
	for _, promStrategy := range promStrategies {
		list = append(list, &PromStrategyItem{ID: promStrategy.ID, Alert: promStrategy.Alert})
	}
	return list
}
