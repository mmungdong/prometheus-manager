package alarmPage

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataAlarmPage "prometheus-manager/cmd/prom_server/internal/data/alarm_page"
)

type (
	// ListReq ...
	ListReq struct {
		// add request params
		Keyword string `form:"keyword"`
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`

		IsDelete bool `form:"is_delete"`
	}

	// ListResp ...
	ListResp struct {
		// add response params
		Page *query.Page `json:"page"`
		List []*Item     `json:"list"`
	}
)

// List ...
func (l *AlarmPage) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	alarmPageData := dataAlarmPage.NewAlarmPage()

	pgInfo := query.NewPage(req.Curr, req.Size)

	alarmList, err := alarmPageData.WithContext(ctx).List(pgInfo)
	if err != nil {
		return nil, err
	}

	list := NewDO(alarmList...).PO().List()

	return &ListResp{
		Page: pgInfo,
		List: list,
	}, nil
}
