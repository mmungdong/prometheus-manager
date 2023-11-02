package alert

import (
	"context"
	"time"

	"gorm.io/gorm/clause"
	dataAlarmHistory "prometheus-manager/cmd/prom_server/internal/data/alarm_history"
	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/model"
)

type (
	// HookReq ...
	HookReq struct {
		// add request params
		*alert.Data
	}

	// HookResp ...
	HookResp struct {
		// add response params
		Msg string `json:"msg"`
	}
)

// PostHook ...
func (l *Alert) PostHook(ctx context.Context, req *HookReq) (*HookResp, error) {
	historyDataOp := dataAlarmHistory.NewAlarmHistory()

	historyDataList := make([]*model.PromAlarmHistory, 0, len(req.Alerts))
	for _, alertItem := range req.Alerts {
		startAt := alert.ParseTime(alertItem.StartsAt).Unix()
		endAt := int64(0)
		duration := time.Now().Unix() - startAt
		if alertItem.EndsAt != "" {
			endAt = alert.ParseTime(alertItem.EndsAt).Unix()
			duration = endAt - startAt
		}

		historyData := &model.PromAlarmHistory{
			Instance:   alertItem.Labels.GetInstance(),
			Status:     req.Status,
			Info:       alertItem.String(),
			StartAt:    alert.ParseTime(alertItem.StartsAt).Unix(),
			EndAt:      alert.ParseTime(alertItem.EndsAt).Unix(),
			Duration:   duration,
			StrategyID: alertItem.Labels.GetStrategyID(),
			LevelID:    0,
			Md5:        alertItem.HashKey(),
			Pages:      nil,
		}
		historyDataList = append(historyDataList, historyData)
	}

	clauses := clause.OnConflict{
		Columns:   []clause.Column{{Name: "md5"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "end_at", "duration"}),
	}
	// 记录告警历史
	if err := historyDataOp.WithContext(ctx).Clauses(clauses).BatchCreate(historyDataList, 100); err != nil {
		return nil, err
	}

	// add your code here
	return &HookResp{}, nil
}
