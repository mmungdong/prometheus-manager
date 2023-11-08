package strategy

import (
	"context"
	"mime/multipart"

	ginplus "github.com/aide-cloud/gin-plus"
	"go.uber.org/zap"
)

type (
	// ImportReq ...
	ImportReq struct {
		// add request params
		RuleYamlFiles []*multipart.FileHeader `form:"rule_yaml_files" binding:"required"`
		GroupId       uint                    `uri:"group_id"`
	}

	// ImportResp ...
	ImportResp struct {
		// add response params
		StrategyCount int `json:"strategy_count"`
	}
)

// PostImport ...
func (l *Strategy) PostImport(ctx context.Context, req *ImportReq) (*ImportResp, error) {
	strategyInfoList, err := l.strategyRepository.ParseStrategyFiles(ctx, req.RuleYamlFiles, req.GroupId)
	if err != nil {
		ginplus.Logger().Error("parse strategy files error", zap.Error(err))
		return nil, err
	}

	if err = l.strategyRepository.BatchCreateStrategy(ctx, strategyInfoList); err != nil {
		return nil, err
	}
	// add your code here
	return &ImportResp{
		StrategyCount: len(strategyInfoList),
	}, nil
}
