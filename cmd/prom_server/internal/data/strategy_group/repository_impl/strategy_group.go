package repository_impl

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/spf13/viper"
	dataStrategyGroup "prometheus-manager/cmd/prom_server/internal/data/strategy_group"
	strategyGroup "prometheus-manager/cmd/prom_server/internal/service/strategy_group"
	"prometheus-manager/pkg/model"
	"prometheus-manager/pkg/strategy"
)

var _ strategyGroup.GroupRepository = (*groupRepositoryImpl)(nil)

type (
	// groupRepositoryImpl ...
	groupRepositoryImpl struct {
		groupData *dataStrategyGroup.StrategyGroup
	}

	// Option ...
	Option func(*groupRepositoryImpl)
)

func (l *groupRepositoryImpl) ExportStrategyGroupFile(_ context.Context, groupDetail *strategy.Group) (*os.File, error) {
	filename := groupDetail.Name + ".yaml"

	vi := viper.New()
	vi.SetConfigType("yaml")
	vi.Set("groups", groupDetail)

	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	if err = vi.WriteConfigAs(filename); err != nil {
		return nil, err
	}

	return f, nil
}

func (l *groupRepositoryImpl) GetStrategyGroupDetailById(ctx context.Context, id uint) (*model.PromGroup, error) {
	fmt.Println("GetStrategyGroupDetailById", l.groupData)
	return l.groupData.WithContext(ctx).Preload(l.groupData.PreloadPromStrategiesKey).FirstByID(id)
}

func (l *groupRepositoryImpl) ParseStrategyGroupFile(ctx context.Context, yamlFiles []*multipart.FileHeader) ([]*model.PromGroup, error) {
	return nil, nil
}

func (l *groupRepositoryImpl) BatchCreateStrategyGroup(ctx context.Context, groups []*model.PromGroup) error {
	if groups == nil || len(groups) == 0 {
		return nil
	}
	return l.groupData.WithContext(ctx).BatchCreate(groups, 100)
}

// defaultGroupRepository ...
func defaultGroupRepository() *groupRepositoryImpl {
	return &groupRepositoryImpl{
		groupData: dataStrategyGroup.NewStrategyGroup(),
	}
}

// NewGroupRepository ...
func NewGroupRepository(opts ...Option) strategyGroup.GroupRepository {
	m := defaultGroupRepository()
	for _, o := range opts {
		o(m)
	}

	return m
}

// WithGroupData ...
func WithGroupData(groupData *dataStrategyGroup.StrategyGroup) Option {
	return func(m *groupRepositoryImpl) {
		m.groupData = groupData
	}
}
