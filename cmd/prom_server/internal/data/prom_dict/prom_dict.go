package dataPromDict

import (
	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"

	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// PromDict mapped from table <prom_dict>
	PromDict struct {
		query.IAction[model.PromDict] `json:"-"`
	}

	// PromDictOption PromDict's option
	PromDictOption func(*PromDict)
)

// NewPromDict new a PromDict
func NewPromDict(options ...PromDictOption) *PromDict {
	promDict := &PromDict{
		IAction: query.NewAction(query.WithDB[model.PromDict](conn.GetMysqlDB())),
	}
	for _, option := range options {
		option(promDict)
	}
	return promDict
}
