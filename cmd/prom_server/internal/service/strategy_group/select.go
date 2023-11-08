package strategyGroup

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
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
func (l *StrategyGroup) GetSelectList(ctx context.Context, req *SelectReq) (*SelectResp, error) {
	strategyGroupData := dataStrategyGroup.NewStrategyGroup()

	var list []*model.PromGroup
	var err error

	pgInfo := query.NewPage(req.Curr, req.Size)
	var wheres []query.ScopeMethod
	if req.Keyword != "" {
		wheres = append(wheres, query.WhereLikeKeyword(req.Keyword+"%", "name"))
	}

	list, err = strategyGroupData.WithContext(ctx).List(pgInfo, wheres...)

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
