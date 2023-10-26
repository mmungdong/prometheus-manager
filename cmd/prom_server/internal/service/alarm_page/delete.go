package alarmPage

import (
	"context"

	dataAlarmPage "prometheus-manager/cmd/prom_server/internal/data/alarm_page"
)

type (
	// DeleteReq ...
	DeleteReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
	}

	// DeleteResp ...
	DeleteResp struct {
		// add response params
		ID uint `json:"id"`
	}
)

// Delete ...
func (l *AlarmPage) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	alarmPageData := dataAlarmPage.NewAlarmPage()

	if err := alarmPageData.WithContext(ctx).DeleteByID(req.ID); err != nil {
		return nil, err
	}
	// add your code here
	return &DeleteResp{ID: req.ID}, nil
}
