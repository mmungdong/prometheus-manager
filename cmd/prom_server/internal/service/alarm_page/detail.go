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
		*Item
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
		Item: NewDO(alarmPageDetail).PO().One(),
	}, nil
}
