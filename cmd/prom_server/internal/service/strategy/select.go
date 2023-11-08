package strategy

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
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
func (l *Strategy) GetSelectList(ctx context.Context, req *SelectReq) (*SelectResp, error) {
	strategyData := dataStrategy.NewStrategy()

	var list []*model.PromStrategy
	var err error

	pgInfo := query.NewPage(req.Curr, req.Size)
	var wheres []query.ScopeMethod
	if req.Keyword != "" {
		wheres = append(wheres, query.WhereLikeKeyword(req.Keyword+"%", "alert"))
	}

	list, err = strategyData.WithContext(ctx).List(pgInfo, wheres...)

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
