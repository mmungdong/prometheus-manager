package alert

import (
	"context"

	"gorm.io/gorm/clause"
	dataAlarmHistory "prometheus-manager/cmd/prom_server/internal/data/alarm_history"
	"prometheus-manager/pkg/alert"
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
	historyDataPO := NewPO(req.Alerts...)
	historyDataDO := historyDataPO.DO()

	clauses := clause.OnConflict{
		Columns:   []clause.Column{{Name: "md5"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "end_at", "duration"}),
	}
	// 记录告警历史
	historyDataOp := dataAlarmHistory.NewAlarmHistory()
	if err := historyDataOp.WithContext(ctx).Clauses(clauses).BatchCreate(historyDataDO.List(), 100); err != nil {
		return nil, err
	}

	// add your code here
	return &HookResp{}, nil
}
