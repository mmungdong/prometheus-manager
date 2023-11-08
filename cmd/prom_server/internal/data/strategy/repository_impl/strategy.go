package repository_impl

import (
	"context"
	"mime/multipart"

	"github.com/spf13/viper"
	dataStrategy "prometheus-manager/cmd/prom_server/internal/data/strategy"
	"prometheus-manager/cmd/prom_server/internal/service/strategy"
	"prometheus-manager/pkg/alert"
	"prometheus-manager/pkg/model"
	promstrategy "prometheus-manager/pkg/strategy"
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

func (l *strategyRepository) ParseStrategyFiles(_ context.Context, yamlFiles []*multipart.FileHeader, groupId uint) ([]*model.PromStrategy, error) {
	var strategyInfoList []*model.PromStrategy
	for _, yamlFile := range yamlFiles {
		strategyGroupInfoList, err := unmarshalStrategyGroup(yamlFile)
		if err != nil {
			return nil, err
		}

		for _, strategyInfo := range strategyGroupInfoList {
			strategyModelInfoList := buildStrategyInfo(strategyInfo, groupId)
			strategyInfoList = append(strategyInfoList, strategyModelInfoList...)
		}
	}

	return strategyInfoList, nil
}

func buildStrategyInfo(strategyInfo *promstrategy.Group, groupId uint) []*model.PromStrategy {
	if strategyInfo == nil {
		return nil
	}

	rules := strategyInfo.Rules
	if rules == nil {
		return nil
	}

	strategyInfoList := make([]*model.PromStrategy, 0, len(rules))
	for _, rule := range rules {
		strategyInfoList = append(strategyInfoList, buildStrategyInfoByRule(rule, groupId))
	}

	return strategyInfoList
}

func buildStrategyInfoByRule(rule *promstrategy.Rule, groupId uint) *model.PromStrategy {
	if rule == nil {
		return nil
	}

	info := &model.PromStrategy{
		GroupID:     groupId,
		Alert:       rule.Alert,
		Expr:        rule.Expr,
		For:         rule.For,
		Labels:      alert.Labels(rule.Labels),
		Annotations: alert.Annotations(rule.Annotations),
		Status:      model.Disable,
	}

	info.ID = uint(info.Labels.GetStrategyID())

	return info
}

// unmarshalStrategyGroup ...
func unmarshalStrategyGroup(yamlFile *multipart.FileHeader) ([]*promstrategy.Group, error) {
	vi := viper.New()
	vi.SetConfigType("yaml")
	file, err := yamlFile.Open()
	defer file.Close()
	if err != nil {
		return nil, err
	}
	if err = vi.ReadConfig(file); err != nil {
		return nil, err
	}

	var strategyGroup []*promstrategy.Group
	if err = vi.UnmarshalKey("groups", &strategyGroup); err != nil {
		return nil, err
	}

	return strategyGroup, nil
}

func (l *strategyRepository) BatchCreateStrategy(ctx context.Context, strategies []*model.PromStrategy) error {
	return l.strategyData.WithContext(ctx).BatchCreate(strategies, 100)
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
