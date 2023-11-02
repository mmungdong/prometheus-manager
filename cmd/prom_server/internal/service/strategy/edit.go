package strategy

import (
	"context"

	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type (
	// EditStrategyReq ...
	EditStrategyReq struct {
		ID uint `uri:"id" binding:"required"`

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

	// EditStrategyResp ...
	EditStrategyResp struct {
		ID uint `json:"id"`
	}
)

func (l *Strategy) Edit(ctx context.Context, req *EditStrategyReq) (*EditStrategyResp, error) {
	strategyData := dataStrategy.NewStrategy()

	firstDetail, err := strategyData.WithContext(ctx).FirstByID(req.ID)
	if err != nil {
		return nil, err
	}

	updateMap := buildUpdateMap(req, firstDetail)

	alarmPages := make([]schema.Tabler, 0, len(req.AlarmPageIds))
	for _, alarmPageId := range req.AlarmPageIds {
		alarmPages = append(alarmPages, &model.PromAlarmPage{BaseModel: query.BaseModel{ID: alarmPageId}})
	}
	promDictList := make([]schema.Tabler, 0, len(req.PromDictIds))
	for _, promDictId := range req.PromDictIds {
		promDictList = append(promDictList, &model.PromDict{BaseModel: query.BaseModel{ID: promDictId}})
	}

	err = strategyData.DB().Transaction(func(tx *gorm.DB) error {
		if err = strategyData.WithDB(tx).WithContext(ctx).UpdateMapByID(req.ID, updateMap); err != nil {
			return err
		}

		// 替换关联
		if err = strategyData.WithDB(tx).WithContext(ctx).Scopes(query.WhereID(req.ID)).AssociationReplace(strategyData.PreloadAlarmPagesKey, alarmPages...); err != nil {
			return err
		}

		if err = strategyData.WithDB(tx).WithContext(ctx).Scopes(query.WhereID(req.ID)).AssociationReplace(strategyData.PreloadCategoriesKey, promDictList...); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &EditStrategyResp{ID: req.ID}, nil
}

func buildUpdateMap(req *EditStrategyReq, firstDetail *model.PromStrategy) map[string]any {
	updateMap := map[string]interface{}{}

	if req.GroupId != firstDetail.GroupID {
		updateMap["group_id"] = req.GroupId
	}
	if req.Alert != firstDetail.Alert {
		updateMap["alert"] = req.Alert
	}
	if req.Expr != firstDetail.Expr {
		updateMap["expr"] = req.Expr
	}
	if req.For != firstDetail.For {
		updateMap["for"] = req.For
	}
	if !req.Labels.Equal(firstDetail.Labels) {
		updateMap["labels"] = req.Labels
	}
	if !req.Annotations.Equal(firstDetail.Annotations) {
		updateMap["annotations"] = req.Annotations
	}

	alarmPages := make([]*model.PromAlarmPage, 0, len(req.AlarmPageIds))
	for _, alarmPageId := range req.AlarmPageIds {
		alarmPages = append(alarmPages, &model.PromAlarmPage{BaseModel: query.BaseModel{ID: alarmPageId}})
	}
	updateMap["alarm_pages"] = alarmPages
	promDictList := make([]*model.PromDict, 0, len(req.PromDictIds))
	for _, promDictId := range req.PromDictIds {
		promDictList = append(promDictList, &model.PromDict{BaseModel: query.BaseModel{ID: promDictId}})
	}
	updateMap["categories"] = promDictList

	return updateMap
}
