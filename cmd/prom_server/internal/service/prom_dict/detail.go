package promDict

import (
	"context"

	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
	"prometheus-manager/pkg/model"
)

type (
	// DetailReq ...
	DetailReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
		// IsDelete 是否删除
		IsDelete bool `form:"is_delete"`
	}

	// DetailResp ...
	DetailResp struct {
		*Item
	}
)

// Detail ...
func (l *PromDict) Detail(ctx context.Context, req *DetailReq) (*DetailResp, error) {
	dictData := dataPromDict.NewPromDict()
	var detailInfo *model.PromDict
	var err error
	if req.IsDelete {
		detailInfo, err = dictData.WithContext(ctx).FirstByIDWithTrashed(req.ID)
	} else {
		detailInfo, err = dictData.WithContext(ctx).FirstByID(req.ID)
	}
	if err != nil {
		return nil, err
	}
	// add your code here
	return &DetailResp{
		Item: NewDO(detailInfo).PO().One(),
	}, nil
}
