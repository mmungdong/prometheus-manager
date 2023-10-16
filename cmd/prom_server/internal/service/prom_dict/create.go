package promDict

import (
	"context"

	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
	"prometheus-manager/pkg/model"
)

type (
	// CreateReq ...
	CreateReq struct {
		// Name 字典名称
		Name string `json:"name" binding:"required"`
		// Category 字典类型
		Category model.Category `json:"category" binding:"required"`
		// Color 展示的颜色, 用于控制tag
		Color string `json:"color"`
		// Remark 备注
		Remark string `json:"remark"`
	}

	// CreateResp ...
	CreateResp struct {
		// add response params
		ID uint `json:"id"`
	}
)

// Create ...
func (l *PromDict) Create(ctx context.Context, req *CreateReq) (*CreateResp, error) {
	dictData := dataPromDict.NewPromDict()

	newDict := &model.PromDict{
		Name:     req.Name,
		Category: req.Category,
		Color:    req.Color,
		Status:   model.Enable,
		Remark:   req.Remark,
	}

	if err := dictData.WithContext(ctx).Create(newDict); err != nil {
		return nil, err
	}
	// add your code here
	return &CreateResp{
		ID: newDict.ID,
	}, nil
}
