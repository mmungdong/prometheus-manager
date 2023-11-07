package alert

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"prometheus-manager/pkg/util/stringer"
)

var (
	_ driver.Valuer = (*KV)(nil)
	_ sql.Scanner   = (*KV)(nil)
	_ fmt.Stringer  = (*Alert)(nil)
)

const (
	// LabelStrategyID 策略ID
	LabelStrategyID = "strategy_id"
	LabelInstance   = "instance"
)

type (
	KV map[string]string

	Labels            = KV
	GroupLabels       = KV
	Annotations       = KV
	CommonLabels      = KV
	CommonAnnotations = KV

	Alert struct {
		Status       string      `json:"status" form:"status"`
		Labels       Labels      `json:"labels" form:"labels"`
		Annotations  Annotations `json:"annotations" form:"annotations"`
		StartsAt     string      `json:"startsAt" form:"startsAt"`
		EndsAt       string      `json:"endsAt" form:"endsAt"`
		GeneratorURL string      `json:"generatorURL" form:"generatorURL"`
		Fingerprint  string      `json:"fingerprint" form:"fingerprint"`
	}

	Data struct {
		Receiver          string            `json:"receiver" form:"receiver"`
		Status            string            `json:"status" form:"status"`
		Alerts            []*Alert          `json:"alerts" form:"alerts"`
		GroupLabels       GroupLabels       `json:"groupLabels" form:"groupLabels"`
		CommonLabels      CommonLabels      `json:"commonLabels" form:"commonLabels"`
		CommonAnnotations CommonAnnotations `json:"commonAnnotations" form:"commonAnnotations"`
		ExternalURL       string            `json:"externalURL" form:"externalURL"`
		Version           string            `json:"version" form:"version"`
		GroupKey          string            `json:"groupKey" form:"groupKey"`
		TruncatedAlerts   int32             `json:"truncatedAlerts" form:"truncatedAlerts"`
	}
)

func (l *Alert) String() string {
	str, _ := json.Marshal(l)
	return string(str)
}

// HashKey 告警的唯一标识
func (l *Alert) HashKey() string {
	return fmt.Sprintf("%d-%s-%s", l.Labels.GetStrategyID(), l.Labels.GetInstance(), l.Fingerprint)
}

func (l *KV) Scan(src any) error {
	return json.Unmarshal(src.([]byte), &l)
}

func (l *KV) Value() (driver.Value, error) {
	if l == nil {
		return "{}", nil
	}

	str, err := json.Marshal(l)
	return string(str), err
}

func (l *Data) Byte() []byte {
	if l == nil {
		return nil
	}
	b, _ := json.Marshal(*l)
	return b
}

func (l *KV) Equal(r KV) bool {
	if len(*l) != len(r) {
		return false
	}

	for k, v := range *l {
		if r[k] != v {
			return false
		}
	}

	return true
}

func (l *KV) Byte() []byte {
	if l == nil {
		return nil
	}
	b, _ := json.Marshal(l)
	return b
}

func (l *KV) String() string {
	return string(l.Byte())
}

// GetStrategyID 获取策略ID, 从标签中获取
func (l *Labels) GetStrategyID() int32 {
	label := *l
	if v, ok := label[LabelStrategyID]; ok {
		return stringer.ParseInt32(v)
	}
	return 0
}

// GetInstance 获取实例, 从标签中获取
func (l *Labels) GetInstance() string {
	label := *l
	return label[LabelInstance]
}
