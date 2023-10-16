package alert

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

var (
	_ driver.Valuer = (*KV)(nil)
	_ sql.Scanner   = (*KV)(nil)
)

type (
	KV map[string]string

	Labels            = KV
	GroupLabels       = KV
	Annotations       = KV
	CommonLabels      = KV
	CommonAnnotations = KV

	Alert struct {
		Status       string      `json:"status"`
		Labels       Labels      `json:"labels"`
		Annotations  Annotations `json:"annotations"`
		StartsAt     int64       `json:"startsAt"`
		EndsAt       int64       `json:"endsAt"`
		GeneratorURL string      `json:"generatorURL"`
		Fingerprint  string      `json:"fingerprint"`
	}

	Data struct {
		Receiver          string            `json:"receiver"`
		Status            string            `json:"status"`
		Alerts            []*Alert          `json:"alerts"`
		GroupLabels       GroupLabels       `json:"groupLabels"`
		CommonLabels      CommonLabels      `json:"commonLabels"`
		CommonAnnotations CommonAnnotations `json:"commonAnnotations"`
		ExternalURL       string            `json:"externalURL"`
		Version           string            `json:"version"`
		GroupKey          string            `json:"groupKey"`
		TruncatedAlerts   int32             `json:"truncatedAlerts"`
	}
)

func (l KV) Scan(src any) error {
	return json.Unmarshal(src.([]byte), &l)
}

func (l KV) Value() (driver.Value, error) {
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
	b, _ := json.Marshal(*l)
	return b
}

func (l *KV) String() string {
	return string(l.Byte())
}
