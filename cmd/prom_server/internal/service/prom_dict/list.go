package promDict

import (
	"context"

	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/times"

	query "github.com/aide-cloud/gorm-normalize"
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
		List []*DictItem `json:"list"`
	}

	DictItem struct {
		ID        uint  `json:"id"`
		CreatedAt int64 `json:"create_at"`
		UpdatedAt int64 `json:"update_at"`
		DeletedAt int64 `json:"delete_at"`

		Name     string         `json:"name"`     // 字典名称
		Color    string         `json:"color"`    // 字典tag颜色
		Remark   string         `json:"remark"`   // 字典备注
		Category model.Category `json:"category"` // 字典类型
		Status   model.Status   `json:"status"`   // 状态
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

	resList := make([]*DictItem, 0, len(list))
	for _, item := range list {
		resList = append(resList, &DictItem{
			ID:        item.ID,
			CreatedAt: times.TimeToUnix(item.CreatedAt),
			UpdatedAt: times.TimeToUnix(item.UpdatedAt),
			DeletedAt: int64(item.DeletedAt),
		})
	}
	// add your code here
	return &ListResp{
		Page: pgInfo,
		List: resList,
	}, nil
}
