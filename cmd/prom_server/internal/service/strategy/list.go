package strategy

import (
	"context"

	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/times"

	query "github.com/aide-cloud/gorm-normalize"
)

type (
	// ListStrategyReq ...
	ListStrategyReq struct {
		Keyword string `form:"keyword"`
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`
	}

	// ListStrategyResp ...
	ListStrategyResp struct {
		List []*Item     `json:"list"`
		Page *query.Page `json:"page"`
	}

	Item struct {
		ID    uint   `json:"id"`
		Alert string `json:"alert"`
		Expr  string `json:"expr"`
		For   string `json:"for"`

		Labels alert.Labels `json:"labels"`
		// Annotations 告警文案
		Annotations alert.Annotations `json:"annotations"`
		// AlertLevel 告警级别
		AlertLevel *AlertLevelItem `json:"alert_level"`
		// GroupInfo 所属规则组
		GroupInfo *GroupItem `json:"group_info"`

		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
		DeletedAt int64 `json:"deleted_at"`
	}
)

// List ...
func (l *Strategy) List(ctx context.Context, req *ListStrategyReq) (*ListStrategyResp, error) {
	strategyData := dataStrategy.NewStrategy()

	wheres := []query.ScopeMethod{
		query.WhereLikeKeyword(req.Keyword, "alert"),
		strategyData.PreloadAlertLevel(),
		strategyData.PreloadGroupInfo(),
	}

	pgInfo := query.NewPage(req.Curr, req.Size)

	strategies, err := strategyData.WithContext(ctx).List(pgInfo, wheres...)
	if err != nil {
		return nil, err
	}

	list := make([]*Item, 0, len(strategies))
	for _, strategy := range strategies {
		list = append(list, &Item{
			ID:          strategy.ID,
			Alert:       strategy.Alert,
			Expr:        strategy.Expr,
			For:         strategy.For,
			Labels:      strategy.Labels,
			Annotations: strategy.Annotations,
			AlertLevel:  buildAlertLevelItem(strategy.AlertLevel),
			GroupInfo:   buildGroupItem(strategy.GroupInfo),
			CreatedAt:   times.TimeToUnix(strategy.CreatedAt),
			UpdatedAt:   times.TimeToUnix(strategy.UpdatedAt),
			DeletedAt:   int64(strategy.DeletedAt),
		})
	}

	return &ListStrategyResp{
		List: list,
		Page: pgInfo,
	}, nil
}
