package promDict

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
	"prometheus-manager/pkg/model"
)

type (
	// ListReq ...
	ListReq struct {
		// add request params
		Keyword string `form:"keyword"`
		Curr    int    `form:"curr"`
		Size    int    `form:"size"`

		// IsDelete 是否删除
		IsDelete bool `form:"is_delete"`
	}

	// ListResp ...
	ListResp struct {
		// add response params
		Page *query.Page `json:"page"`
		List []*Item     `json:"list"`
	}
)

// List ...
func (l *PromDict) List(ctx context.Context, req *ListReq) (*ListResp, error) {
	dictData := dataPromDict.NewPromDict()

	var list []*model.PromDict
	var err error
	pgInfo := query.NewPage(req.Curr, req.Size)

	wheres := []query.ScopeMethod{query.WhereLikeKeyword(req.Keyword, "name")}

	if req.IsDelete {
		list, err = dictData.WithContext(ctx).ListWithTrashed(pgInfo, wheres...)
	} else {
		list, err = dictData.WithContext(ctx).List(pgInfo, wheres...)
	}

	if err != nil {
		return nil, err
	}

	resList := NewDO(list...).PO().List()

	// add your code here
	return &ListResp{
		Page: pgInfo,
		List: resList,
	}, nil
}
