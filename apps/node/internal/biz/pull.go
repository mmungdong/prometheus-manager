package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "prometheus-manager/api/strategy/v1/pull"
	"prometheus-manager/apps/node/internal/service"
)

type (
	IPullRepo interface {
		V1Repo
	}

	PullLogic struct {
		logger *log.Helper
		repo   IPullRepo
	}
)

var _ service.IPullLogic = (*PullLogic)(nil)

func NewPullLogic(repo IPullRepo, logger log.Logger) *PullLogic {
	return &PullLogic{repo: repo, logger: log.NewHelper(log.With(logger, "module", "biz/Pull"))}
}

func (s *PullLogic) Strategies(ctx context.Context, req *pb.StrategiesRequest) (*pb.StrategiesReply, error) {
	return nil, nil
}
