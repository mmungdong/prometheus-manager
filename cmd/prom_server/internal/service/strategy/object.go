package strategy

import (
	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/plugin/soft_delete"
	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/object"
	"prometheus-manager/pkg/times"
)

type (
	Item struct {
		ID    uint   `json:"id"`
		Alert string `json:"alert"`
		Expr  string `json:"expr"`
		For   string `json:"for"`

		Labels alert.Labels `json:"labels"`
		// Annotations 告警文案
		Annotations alert.Annotations `json:"annotations"`
		// AlertLevel 告警级别
		AlertLevel *AlertLevelItem `json:"alert_level"`
		// GroupInfo 所属规则组
		GroupInfo  *GroupItem       `json:"group_info"`
		AlarmPages []*AlarmPageItem `json:"alarm_pages"`
		Categories []*CategoryItem  `json:"categories"`
		Status     model.Status     `json:"status"`

		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
		DeletedAt int64 `json:"deleted_at"`

		// GroupId 所属规则组ID
		GroupId uint `json:"group_id"`
		// AlarmPageIds 告警页面
		AlarmPageIds []uint `json:"alarm_page_ids"`
		// PromDictIds 规则类型
		PromDictIds []uint `json:"prom_dict_ids"`
		// AlertLevelId 告警级别
		AlertLevelId uint `json:"alert_level_id"`
	}

	GroupItem struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	AlarmPageItem struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	CategoryItem struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	AlertLevelItem struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)

// NewDO 实例化strategy DO
func NewDO(values ...*model.PromStrategy) *object.DO[model.PromStrategy, Item] {
	return object.NewDO(
		object.DOWithList[model.PromStrategy, Item](values...),
		object.DOWithBuildFunc[model.PromStrategy, Item](modelToStrategy),
	)
}

// NewPO 实例化strategy PO
func NewPO(values ...*Item) *object.PO[model.PromStrategy, Item] {
	return object.NewPO(
		object.POWithList[model.PromStrategy, Item](values...),
		object.POWithBuildFunc[model.PromStrategy, Item](strategyToModel),
	)
}

// strategyToModel ...
func strategyToModel(strategy *Item) *model.PromStrategy {
	alarmPages := make([]*model.PromAlarmPage, 0, len(strategy.AlarmPageIds))
	for _, alarmPageId := range strategy.AlarmPageIds {
		alarmPages = append(alarmPages, &model.PromAlarmPage{BaseModel: query.BaseModel{ID: alarmPageId}})
	}
	promDictList := make([]*model.PromDict, 0, len(strategy.PromDictIds))
	for _, promDictId := range strategy.PromDictIds {
		promDictList = append(promDictList, &model.PromDict{BaseModel: query.BaseModel{ID: promDictId}})
	}
	return &model.PromStrategy{
		Alert:        strategy.Alert,
		Expr:         strategy.Expr,
		For:          strategy.For,
		Labels:       strategy.Labels,
		Annotations:  strategy.Annotations,
		GroupID:      strategy.GroupId,
		AlertLevelID: strategy.AlertLevelId,
		AlarmPages:   alarmPages,
		Categories:   promDictList,
		BaseModel: query.BaseModel{
			ID:        strategy.ID,
			CreatedAt: times.UnixToTime(strategy.CreatedAt),
			UpdatedAt: times.UnixToTime(strategy.UpdatedAt),
			DeletedAt: soft_delete.DeletedAt(strategy.DeletedAt),
		},
	}
}

// modelToStrategy ...
func modelToStrategy(modelItem *model.PromStrategy) *Item {
	return &Item{
		ID:          modelItem.ID,
		Alert:       modelItem.Alert,
		Expr:        modelItem.Expr,
		For:         modelItem.For,
		Labels:      modelItem.Labels,
		Annotations: modelItem.Annotations,
		AlertLevel:  buildAlertLevelItem(modelItem.AlertLevel),
		GroupInfo:   buildGroupItem(modelItem.GroupInfo),
		CreatedAt:   times.TimeToUnix(modelItem.CreatedAt),
		UpdatedAt:   times.TimeToUnix(modelItem.UpdatedAt),
		DeletedAt:   int64(modelItem.DeletedAt),
		AlarmPages:  buildAlarmPageItems(modelItem.AlarmPages),
		Categories:  buildCategoryItems(modelItem.Categories),
		Status:      modelItem.Status,
		GroupId:     modelItem.GroupID,
	}
}

func buildGroupItem(group *model.PromGroup) *GroupItem {
	if group == nil {
		return nil
	}

	return &GroupItem{
		ID:   group.ID,
		Name: group.Name,
	}
}

func buildAlertLevelItem(alertLevel *model.PromDict) *AlertLevelItem {
	if alertLevel == nil {
		return nil
	}

	return &AlertLevelItem{
		ID:   alertLevel.ID,
		Name: alertLevel.Name,
	}
}

func buildAlarmPageItems(alarmPages []*model.PromAlarmPage) []*AlarmPageItem {
	alarmPageItems := make([]*AlarmPageItem, 0, len(alarmPages))

	for _, alarmPage := range alarmPages {
		if alarmPage == nil {
			continue
		}
		alarmPageItems = append(alarmPageItems, &AlarmPageItem{
			ID:   alarmPage.ID,
			Name: alarmPage.Name,
		})
	}

	return alarmPageItems
}

func buildCategoryItems(categories []*model.PromDict) []*CategoryItem {
	categoryItems := make([]*CategoryItem, 0, len(categories))

	for _, category := range categories {
		if category == nil {
			continue
		}
		categoryItems = append(categoryItems, &CategoryItem{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return categoryItems
}
