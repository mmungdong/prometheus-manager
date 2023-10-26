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

	// Item ...
	Item struct {
		ID            uint         `json:"id"`
		Name          string       `json:"name"`
		Remark        string       `json:"remark"`
		StrategyCount int64        `json:"strategy_count"`
		Status        model.Status `json:"status"`
	}
)

// PostList ...
func (l *StrategyGroup) PostList(ctx context.Context, req *ListReq) (*ListResp, error) {
	strategyGrouData := dataStrategyGroup.NewStrategyGroup()

	pgInfo := query.NewPage(req.Curr, req.Size)
	wheres := []query.Scopemethod{query.WhereLikeKeyword(req.Keyword, "name")}

	var err error
	var strategyGroups []*model.PromGroup

	// 查询
	if req.IsDelete {
		strategyGroups, err = strategyGrouData.WithContext(ctx).ListWithTrashed(pgInfo, wheres...)
	} else {
		strategyGroups, err = strategyGrouData.WithContext(ctx).List(pgInfo, wheres...)
	}

	if err != nil {
		return nil, err
	}

	list := make([]*Item, 0, len(strategyGroups))
	for _, strategyGroup := range strategyGroups {
		list = append(list, &Item{
			ID:            strategyGroup.ID,
			Name:          strategyGroup.Name,
			Remark:        strategyGroup.Remark,
			StrategyCount: strategyGroup.StrategyCount,
			Status:        strategyGroup.Status,
		})
	}
	// add your code here
	return &ListResp{
		Page: pgInfo,
		List: list,
	}, nil
}
