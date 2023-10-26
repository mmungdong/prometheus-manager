package alarmPage

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"

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

	promStrategies := make([]*model.PromStrategy, 0, len(req.PromStrategyIds))
	for _, id := range req.PromStrategyIds {
		promStrategies = append(promStrategies, &model.PromStrategy{
			BaseModel: query.BaseModel{ID: id},
		})
	}

	newAlarmPage := &model.PromAlarmPage{
		Name:           req.Name,
		Remark:         req.Remark,
		Icon:           req.Icon,
		Color:          req.Color,
		Status:         req.Status,
		PromStrategies: promStrategies,
	}

	if err := historyData.WithContext(ctx).Create(newAlarmPage); err != nil {
		return nil, err
	}
	// add your code here
	return &CreateResp{
		ID: newAlarmPage.ID,
	}, nil
}
