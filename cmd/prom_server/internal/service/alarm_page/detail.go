package alarmPage

import (
	"context"
	dataAlarmPage "prometheus-manager/cmd/prom_server/internal/data/alarm_page"
	"prometheus-manager/pkg/model"
)

type (
	// DetailReq ...
	DetailReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
		// IsDelete 是否查询已删除的数据
		IsDelete bool `form:"is_delete"`
	}

	// DetailResp ...
	DetailResp struct {
		ID             uint                `json:"id"`
		Name           string              `json:"name"`
		Remark         string              `json:"remark"`
		Icon           string              `json:"icon"`
		Color          string              `json:"color"`
		Status         model.Status        `json:"status"`
		PromStrategies []*PromStrategyItem `json:"prom_strategies"`
	}

	PromStrategyItem struct {
		ID    uint   `json:"id"`
		Alert string `json:"alert"`
	}
)

// Detail ...
func (l *AlarmPage) Detail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	alarmPageData := dataAlarmPage.NewAlarmPage()
	var alarmPageDetail *model.PromAlarmPage
	var err error

	if req.IsDelete {
		alarmPageDetail, err = alarmPageData.WithContext(ctx).FirstByIDWithTrashed(req.ID)
	} else {
		alarmPageDetail, err = alarmPageData.WithContext(ctx).FirstByID(req.ID)
	}

	if err != nil {
		return nil, err
	}

	return &DetailResp{
		ID:             alarmPageDetail.ID,
		Name:           alarmPageDetail.Name,
		Remark:         alarmPageDetail.Remark,
		Icon:           alarmPageDetail.Icon,
		Color:          alarmPageDetail.Color,
		Status:         alarmPageDetail.Status,
		PromStrategies: ToPromStrategyItems(alarmPageDetail.PromStrategies),
	}, nil
}

// ToPromStrategyItems ...
func ToPromStrategyItems(list []*model.PromStrategy) []*PromStrategyItem {
	items := make([]*PromStrategyItem, 0, len(list))
	for _, item := range list {
		items = append(items, &PromStrategyItem{
			ID:    item.ID,
			Alert: item.Alert,
		})
	}
	return items
}
