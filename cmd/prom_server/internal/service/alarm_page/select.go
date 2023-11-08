package alarmPage

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataAlarmPage "prometheus-manager/cmd/prom_server/internal/data/alarm_page"
	"prometheus-manager/pkg/model"
)

type (
	// SelectReq ...
	SelectReq struct {
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`
		Keyword string `form:"keyword"`
	}

	// SelectResp ...
	SelectResp struct {
		List []*SelectItem `json:"list"`
		Page *query.Page   `json:"page"`
	}
)

// GetSelectList ...
func (l *AlarmPage) GetSelectList(ctx context.Context, req *SelectReq) (*SelectResp, error) {
	alarmPageData := dataAlarmPage.NewAlarmPage()

	var list []*model.PromAlarmPage
	var err error

	pgInfo := query.NewPage(req.Curr, req.Size)
	wheres := []query.ScopeMethod{
		query.WhereLikeKeyword(req.Keyword+"%", "name"),
	}

	list, err = alarmPageData.WithContext(ctx).List(pgInfo, wheres...)

	if err != nil {
		return nil, err
	}

	resList := NewSelectDO(list...).PO().List()

	// add your code here
	return &SelectResp{
		List: resList,
		Page: pgInfo,
	}, nil
}
