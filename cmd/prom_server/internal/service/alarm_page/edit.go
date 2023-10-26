package alarmPage

import (
	"context"
	
	dataAlarmPage "prometheus-manager/cmd/prom_server/internal/data/alarm_page"
	"prometheus-manager/pkg/model"
)

type (
	// EditReq ...
	EditReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`

		// Name 报警页面名称
		Name   string       `json:"name" binding:"required"`
		Remark string       `json:"remark" binding:"required"`
		Icon   string       `json:"icon" binding:"required"`
		Color  string       `json:"color" binding:"required"`
		Status model.Status `json:"status" binding:"required"`
	}

	// EditResp ...
	EditResp struct {
		// add response params
		ID uint `json:"id"`
	}
)

// Edit ...
func (l *AlarmPage) Edit(ctx context.Context, req *EditReq) (*EditResp, error) {
	alarmPageData := dataAlarmPage.NewAlarmPage()

	alarmPageDetail, err := alarmPageData.WithContext(ctx).FirstByID(req.ID)
	if err != nil {
		return nil, err
	}

	updateMap := l.getUpdateMap(req, alarmPageDetail)

	if len(updateMap) > 0 {
		if err := alarmPageData.WithContext(ctx).UpdateMapByID(req.ID, updateMap); err != nil {
			return nil, err
		}
	}
	// add your code here
	return &EditResp{ID: req.ID}, nil
}

func (l *AlarmPage) getUpdateMap(req *EditReq, alarmPageDetail *model.PromAlarmPage) map[string]any {
	updateMap := make(map[string]any)
	// 对比req和alarmPageDetail的字段，如果不同则更新
	if alarmPageDetail.Name != req.Name {
		updateMap["name"] = req.Name
	}

	if alarmPageDetail.Remark != req.Remark {
		updateMap["remark"] = req.Remark
	}

	if alarmPageDetail.Icon != req.Icon {
		updateMap["icon"] = req.Icon
	}

	if alarmPageDetail.Color != req.Color {
		updateMap["color"] = req.Color
	}

	if alarmPageDetail.Status != req.Status {
		updateMap["status"] = req.Status
	}
	return updateMap
}
