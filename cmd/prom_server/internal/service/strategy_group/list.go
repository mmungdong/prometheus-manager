package strategyGroup

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"

	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
	"prometheus-manager/pkg/model"
)

type (
	// ListReq ...
	ListReq struct {
		Keyword string `form:"keyword"`
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`

		IsDelete bool `form:"is_delete"`
	}

	// ListResp ...
	ListResp struct {
		Page *query.Page `json:"page"`
		List []*Item     `json:"list"`
	}
)

// PostList ...
func (l *StrategyGroup) PostList(ctx context.Context, req *ListReq) (*ListResp, error) {
	strategyGroupData := dataStrategyGroup.NewStrategyGroup()

	pgInfo := query.NewPage(req.Curr, req.Size)
	wheres := []query.ScopeMethod{query.WhereLikeKeyword(req.Keyword, "name")}

	var err error
	var strategyGroups []*model.PromGroup

	// 查询
	if req.IsDelete {
		strategyGroups, err = strategyGroupData.WithContext(ctx).ListWithTrashed(pgInfo, wheres...)
	} else {
		strategyGroups, err = strategyGroupData.WithContext(ctx).List(pgInfo, wheres...)
	}

	if err != nil {
		return nil, err
	}

	list := NewDO(strategyGroups...).PO().List()
	// add your code here
	return &ListResp{
		Page: pgInfo,
		List: list,
	}, nil
}
