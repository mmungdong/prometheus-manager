package promDict

import (
	"context"
	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
)

type (
	// EditReq ...
	EditReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`

		// Name 字典名称
		Name string `json:"name" binding:"required"`
		// Remark 备注
		Remark string `json:"remark"`
		// Color 展示的颜色, 用于控制tag
		Color string `json:"color"`
	}

	// EditResp ...
	EditResp struct {
		// add response params
		ID uint `json:"id"`
	}
)

// Edit ...
func (l *PromDict) Edit(ctx context.Context, req *EditReq) (*EditResp, error) {
	dictData := dataPromDict.NewPromDict()

	if err := dictData.WithContext(ctx).UpdateMapByID(req.ID, map[string]any{
		"name":   req.Name,
		"remark": req.Remark,
		"color":  req.Color,
	}); err != nil {
		return nil, err
	}
	// add your code here
	return &EditResp{req.ID}, nil
}
