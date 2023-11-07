package alarmPage

import (
	"context"

	dataAlarmPage "prometheus-manager/cmd/prom_server/internal/data/alarm_page"
	"prometheus-manager/pkg/model"
)

type (
	// CreateReq ...
	CreateReq struct {
		Name            string       `json:"name" binding:"required"`
		Remark          string       `json:"remark"`
		Icon            string       `json:"icon"`
		Color           string       `json:"color"`
		Status          model.Status `json:"status"`
		PromStrategyIds []uint       `json:"prom_strategy_ids"`
	}

	// CreateResp ...
	CreateResp struct {
		ID uint `json:"id"`
	}
)

// Create ...
func (l *AlarmPage) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	historyData := dataAlarmPage.NewAlarmPage()

	alarmPagePOValue := &Item{
		Name:            req.Name,
		Remark:          req.Remark,
		Icon:            req.Icon,
		Color:           req.Color,
		Status:          req.Status,
		promStrategyIds: req.PromStrategyIds,
	}

	newAlarmPage := NewPO(alarmPagePOValue).DO().One()

	if err := historyData.WithContext(ctx).Create(newAlarmPage); err != nil {
		return nil, err
	}
	// add your code here
	return &CreateResp{
		ID: newAlarmPage.ID,
	}, nil
}
