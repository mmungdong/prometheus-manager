package dataPromDict

import (
	query "github.com/aide-cloud/gorm-normalize"

	"prometheus-manager/pkg/conn"
	"prometheus-manager/pkg/model"
)

type (
	// PromDict mapped from table <prom_dict>
	PromDict struct {
		query.IAction[model.PromDict] `json:"-"`
	}

	// PromDictOption PromDict option
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
