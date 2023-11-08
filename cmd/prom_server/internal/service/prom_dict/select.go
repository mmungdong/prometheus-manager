package promDict

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
	"prometheus-manager/pkg/model"
)

type (
	// SelectReq ...
	SelectReq struct {
		Category model.Category `form:"category"`
		Curr     int            `form:"curr"`
		Size     int            `form:"size"`
		Keyword  string         `form:"keyword"`
	}

	// SelectResp ...
	SelectResp struct {
		List []*SelectItem `json:"list"`
		Page *query.Page   `json:"page"`
	}
)

// GetSelectList ...
func (l *PromDict) GetSelectList(ctx context.Context, req *SelectReq) (*SelectResp, error) {
	dictData := dataPromDict.NewPromDict()

	var list []*model.PromDict
	var err error

	pgInfo := query.NewPage(req.Curr, req.Size)
	wheres := []query.ScopeMethod{
		query.WhereLikeKeyword(req.Keyword+"%", "name"),
		dataPromDict.WhereCategory(req.Category),
	}

	list, err = dictData.WithContext(ctx).List(pgInfo, wheres...)

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
