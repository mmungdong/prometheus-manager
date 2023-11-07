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
)

// NewDO 实例化alarmPage DO
func NewDO(values ...*model.PromAlarmPage) *object.DO[model.PromAlarmPage, Item] {
	return object.NewDO(
		object.DOWithList[model.PromAlarmPage, Item](values...),
		object.DOWithBuildFunc[model.PromAlarmPage, Item](modelToAlarmPage),
	)
}

// NewPO 实例化alarmPage PO
func NewPO(values ...*Item) *object.PO[model.PromAlarmPage, Item] {
	return object.NewPO(
		object.POWithList[model.PromAlarmPage, Item](values...),
		object.POWithBuildFunc[model.PromAlarmPage, Item](alarmPageToModel),
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
