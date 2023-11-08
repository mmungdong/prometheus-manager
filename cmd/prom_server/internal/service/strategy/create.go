package strategy

import (
	"context"

	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/pkg/alert"
)

type (
	// CreateStrategyReq ...
	CreateStrategyReq struct {
		GroupId     uint              `json:"group_id" desc:"分组id"`
		Alert       string            `json:"alert" desc:"告警名称"`
		Expr        string            `json:"expr" desc:"表达式"`
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

	newStrategyPOValue := &Item{
		GroupId:      req.GroupId,
		Alert:        req.Alert,
		Expr:         req.Expr,
		For:          req.For,
		Labels:       req.Labels,
		Annotations:  req.Annotations,
		AlarmPageIds: req.AlarmPageIds,
		PromDictIds:  req.PromDictIds,
		AlertLevelId: req.AlertLevelId,
	}

	newStrategy := NewPO(newStrategyPOValue).DO().One()

	if err := strategyData.WithContext(ctx).Create(newStrategy); err != nil {
		return nil, err
	}
	// your code ...
	return &CreateStrategyResp{ID: newStrategy.ID}, nil
}
