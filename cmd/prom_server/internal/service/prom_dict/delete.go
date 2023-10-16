package promDict

import (
	"context"
	dataPromDict "prometheus-manager/cmd/prom_server/internal/data/prom_dict"
)

type (
	// DeleteReq ...
	DeleteReq struct {
		// add request params
		ID uint `uri:"id" binding:"required"`
	}

	// DeleteResp ...
	DeleteResp struct {
		// add response params
		ID uint `json:"id"`
	}
)

// Delete ...
func (l *PromDict) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	dictData := dataPromDict.NewPromDict()

	if err := dictData.WithContext(ctx).DeleteByID(req.ID); err != nil {
		return nil, err
	}
	// add your code here
	return &DeleteResp{ID: req.ID}, nil
}
