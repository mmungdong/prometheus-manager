package strategyGroup

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
		GroupCount    int      `json:"group_count"`
		StrategyCount int      `json:"strategy_count"`
		SuccessTotal  int      `json:"success_total"`
		FailTotal     int      `json:"fail_total"`
		FailList      []string `json:"fail_list"`
	}
)

// PostImport ...
func (l *StrategyGroup) PostImport(ctx context.Context, req *ImportReq) (*ImportResp, error) {
	groupInfoList, err := l.groupRepository.ParseStrategyGroupFile(ctx, req.RuleYamlFiles)
	if err != nil {
		return nil, err
	}

	if err = l.groupRepository.BatchCreateStrategyGroup(ctx, groupInfoList); err != nil {
		return nil, err
	}
	// add your code here
	return &ImportResp{}, nil
}
