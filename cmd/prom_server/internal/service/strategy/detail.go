package strategy

import (
	"context"

	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/times"
)

type (
	// DetailStrategyReq ...
	DetailStrategyReq struct {
		ID uint `uri:"id" binding:"required"`

		IsDelete bool `form:"is_delete"`
	}

	// DetailStrategyResp ...
	DetailStrategyResp struct {
		Group       *GroupItem        `json:"group"`
		ID          uint              `json:"id"`
		Alert       string            `json:"alert"`
		Expr        string            `json:"expr"`
		For         string            `json:"for"`
		Labels      alert.Labels      `json:"labels"`      // 标签
		Annotations alert.Annotations `json:"annotations"` // 告警文案
		AlarmPages  []*AlarmPageItem  `json:"alarm_pages"`
		Categories  []*CategoryItem   `json:"categories"`
		AlertLevel  *AlertLevelItem   `json:"alert_level"`
		Status      model.Status      `json:"status"`

		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
		DeletedAt int64 `json:"deleted_at"`
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

func (l *Strategy) Detail(ctx context.Context, req *DetailStrategyReq) (*DetailStrategyResp, error) {
	strategyData := dataStrategy.NewStrategy()

	var detailInfo *model.PromStrategy
	var err error
	if req.IsDelete {
		detailInfo, err = strategyData.WithContext(ctx).FirstByIDWithTrashed(req.ID)
	} else {
		detailInfo, err = strategyData.WithContext(ctx).FirstByID(req.ID)
	}

	if err != nil {
		return nil, err
	}

	return &DetailStrategyResp{
		ID:          detailInfo.ID,
		Alert:       detailInfo.Alert,
		Expr:        detailInfo.Expr,
		For:         detailInfo.For,
		Labels:      detailInfo.Labels,
		Annotations: detailInfo.Annotations,
		AlertLevel:  buildAlertLevelItem(detailInfo.AlertLevel),
		AlarmPages:  buildAlarmPageItems(detailInfo.AlarmPages),
		Categories:  buildCategoryItems(detailInfo.Categories),
		Group:       buildGroupItem(detailInfo.GroupInfo),
		Status:      detailInfo.Status,
		CreatedAt:   times.TimeToUnix(detailInfo.CreatedAt),
		UpdatedAt:   times.TimeToUnix(detailInfo.UpdatedAt),
		DeletedAt:   int64(detailInfo.DeletedAt),
	}, nil
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
