package strategy

import (
	"context"
	"mime/multipart"
)

type (
	// ImportReq ...
	ImportReq struct {
		// add request params
		RuleYamlFiles []*multipart.FileHeader `form:"rule_yaml_file" binding:"required"`
	}

	// ImportResp ...
	ImportResp struct {
		// add response params
		GroupCount int `json:"group_count"`
	}
)

// PostImport ...
func (l *Strategy) PostImport(ctx context.Context, req *ImportReq) (*ImportResp, error) {
	strategyInfoList, err := l.strategyRepository.ParseStrategyFiles(ctx, req.RuleYamlFiles)
	if err != nil {
		return nil, err
	}

	if err = l.strategyRepository.BatchCreateStrategy(ctx, strategyInfoList); err != nil {
		return nil, err
	}
	// add your code here
	return &ImportResp{}, nil
}
