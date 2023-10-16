package strategy

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/model"
)

type (
	// CreateStrategyReq ...
	CreateStrategyReq struct {
		GroupId     uint              `json:"group_id"`
		Alert       string            `json:"alert"`
		Expr        string            `json:"expr"`
		For         string            `json:"for"`
		Labels      alert.Labels      `json:"labels"`
		Annotations alert.Annotations `json:"annotations"`
		// AlarmPageIds 告警页面
		AlarmPageIds []uint `json:"alarm_page_ids"`
		// PromDictIds 规则类型
		PromDictIds []uint `json:"prom_dict_ids"`
		// AlertLevelId 告警级别
		AlertLevelId uint `json:"alert_level_id"`
	}

	// CreateStrategyResp ...
	CreateStrategyResp struct {
		ID uint `json:"id"`
	}
)

func (l *Strategy) Create(ctx context.Context, req *CreateStrategyReq) (*CreateStrategyResp, error) {
	strategyData := dataStrategy.NewStrategy()

	alarmPages := make([]*model.PromAlarmPage, 0, len(req.AlarmPageIds))
	for _, alarmPageId := range req.AlarmPageIds {
		alarmPages = append(alarmPages, &model.PromAlarmPage{BaseModel: query.BaseModel{ID: alarmPageId}})
	}
	promDicts := make([]*model.PromDict, 0, len(req.PromDictIds))
	for _, promDictId := range req.PromDictIds {
		promDicts = append(promDicts, &model.PromDict{BaseModel: query.BaseModel{ID: promDictId}})
	}

	newStrategy := &model.PromStrategy{
		GroupID:      req.GroupId,
		Alert:        req.Alert,
		Expr:         req.Expr,
		For:          req.For,
		Labels:       req.Labels,
		Annotations:  req.Annotations,
		AlarmPages:   alarmPages,
		Categories:   promDicts,
		AlertLevelID: req.AlertLevelId,
	}

	if err := strategyData.WithContext(ctx).Create(newStrategy); err != nil {
		return nil, err
	}
	// your code ...
	return &CreateStrategyResp{ID: newStrategy.ID}, nil
}
