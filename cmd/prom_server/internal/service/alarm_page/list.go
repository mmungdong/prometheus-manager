package alarmPage

import (
	"context"

	dataAlarmPage "prometheus-manager/cmd/prom_server/internal/data/alarm_page"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/times"

	query "github.com/aide-cloud/gorm-normalize"
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
		Page *query.Page      `json:"page"`
		List []*AlarmPageItem `json:"list"`
	}

	AlarmPageItem struct {
		ID     uint         `json:"id"`
		Name   string       `json:"name"`
		Icon   string       `json:"icon"`
		Color  string       `json:"color"`
		Remark string       `json:"remark"`
		Status model.Status `json:"status"`

		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
		DeletedAt int64 `json:"deleted_at"`
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

	list := make([]*AlarmPageItem, 0, len(alarmList))
	for _, alarmPage := range alarmList {
		list = append(list, &AlarmPageItem{
			ID:        alarmPage.ID,
			Name:      alarmPage.Name,
			Icon:      alarmPage.Icon,
			Color:     alarmPage.Color,
			Remark:    alarmPage.Remark,
			Status:    alarmPage.Status,
			DeletedAt: int64(alarmPage.DeletedAt),
			CreatedAt: times.TimeToUnix(alarmPage.CreatedAt),
			UpdatedAt: times.TimeToUnix(alarmPage.UpdatedAt),
		})
	}

	return &ListResp{
		Page: pgInfo,
		List: list,
	}, nil
}
