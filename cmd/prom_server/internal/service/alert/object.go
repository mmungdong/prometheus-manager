package alert

import (
	"encoding/json"
	"time"

	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/object"
)

// NewDO 实例化alarm_history DO
func NewDO(values ...*model.PromAlarmHistory) *object.DO[model.PromAlarmHistory, alert.Alert] {
	return object.NewDO[model.PromAlarmHistory, alert.Alert](
		object.DOWithList[model.PromAlarmHistory, alert.Alert](values...),
		object.DOWithBuildFunc[model.PromAlarmHistory, alert.Alert](modelToAlertPromAlarmHistory),
	)
}

// NewPO 实例化alarm_history PO
func NewPO(values ...*alert.Alert) *object.PO[model.PromAlarmHistory, alert.Alert] {
	return object.NewPO[model.PromAlarmHistory, alert.Alert](
		object.POWithList[model.PromAlarmHistory, alert.Alert](values...),
		object.POWithBuildFunc[model.PromAlarmHistory, alert.Alert](alertToModelPromAlarmHistory),
	)
}

// alertToModelPromAlarmHistory ...
func alertToModelPromAlarmHistory(alertItem *alert.Alert) *model.PromAlarmHistory {
	startAt := alert.ParseTime(alertItem.StartsAt).Unix()
	var endAt int64
	duration := time.Now().Unix() - startAt
	if alertItem.EndsAt != "" {
		endAt = alert.ParseTime(alertItem.EndsAt).Unix()
		duration = endAt - startAt
	}
	return &model.PromAlarmHistory{
		Instance:   alertItem.Labels.GetInstance(),
		Status:     alertItem.Status,
		Info:       alertItem.String(),
		StartAt:    alert.ParseTime(alertItem.StartsAt).Unix(),
		EndAt:      alert.ParseTime(alertItem.EndsAt).Unix(),
		Duration:   duration,
		StrategyID: alertItem.Labels.GetStrategyID(),
		LevelID:    0,
		Md5:        alertItem.HashKey(),
		Pages:      nil,
	}
}

// modelToAlertPromAlarmHistory ...
func modelToAlertPromAlarmHistory(modelItem *model.PromAlarmHistory) *alert.Alert {
	var alertInfo alert.Alert
	_ = json.Unmarshal([]byte(modelItem.Info), &alertInfo)
	return &alertInfo
}
