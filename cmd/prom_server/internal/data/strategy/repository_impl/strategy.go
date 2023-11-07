package repository_impl

import (
	"context"
	"mime/multipart"

	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/cmd/prom_server/internal/service/strategy"
	"prometheus-manager/pkg/model"
)

var _ strategy.Repository = (*strategyRepository)(nil)

type (
	// strategyRepository ...
	strategyRepository struct {
		strategyData *dataStrategy.Strategy
	}

	// Option ...
	Option func(*strategyRepository)
)

func (l *strategyRepository) ParseStrategyFiles(ctx context.Context, yamlFiles []*multipart.FileHeader) ([]*model.PromStrategy, error) {
	//TODO implement me
	panic("implement me")
}

func (l *strategyRepository) BatchCreateStrategy(ctx context.Context, strategies []*model.PromStrategy) error {
	//TODO implement me
	panic("implement me")
}

// defaultStrategyRepository ...
func defaultStrategyRepository() *strategyRepository {
	return &strategyRepository{
		strategyData: dataStrategy.NewStrategy(),
	}
}

// NewStrategyRepository ...
func NewStrategyRepository(opts ...Option) strategy.Repository {
	m := defaultStrategyRepository()
	for _, o := range opts {
		o(m)
	}

	return m
}

// WithStrategyData ...
func WithStrategyData(data *dataStrategy.Strategy) Option {
	return func(m *strategyRepository) {
		m.strategyData = data
	}
}
