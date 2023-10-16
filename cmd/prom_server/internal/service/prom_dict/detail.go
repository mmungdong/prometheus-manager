package promDict

import (
	"context"
	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/times"
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
		ID        uint  `json:"id"`
		CreatedAt int64 `json:"create_at"`
		UpdatedAt int64 `json:"update_at"`
		DeletedAt int64 `json:"delete_at"`

		Name     string         `json:"name"`     // 字典名称
		Category model.Category `json:"category"` // 字典类型
		Color    string         `json:"color"`    // 字典tag颜色
		Status   model.Status   `json:"status"`   // 状态
		Remark   string         `json:"remark"`   // 字典备注
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
		ID:        detailInfo.ID,
		CreatedAt: times.TimeToUnix(detailInfo.CreatedAt),
		UpdatedAt: times.TimeToUnix(detailInfo.UpdatedAt),
		DeletedAt: int64(detailInfo.DeletedAt),
		Name:      detailInfo.Name,
		Category:  detailInfo.Category,
		Color:     detailInfo.Color,
		Status:    detailInfo.Status,
		Remark:    detailInfo.Remark,
	}, nil
}
