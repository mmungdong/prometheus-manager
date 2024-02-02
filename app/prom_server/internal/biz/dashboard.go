package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"prometheus-manager/app/prom_server/internal/biz/bo"
	"prometheus-manager/app/prom_server/internal/biz/do/basescopes"
	"prometheus-manager/app/prom_server/internal/biz/repository"
)

type (
	DashboardBiz struct {
		log *log.Helper

		dashboardRepo repository.DashboardRepo
		chartRepo     repository.ChartRepo
	}
)

func NewDashboardBiz(
	dashboardRepo repository.DashboardRepo,
	chartRepo repository.ChartRepo,
	logger log.Logger,
) *DashboardBiz {
	return &DashboardBiz{
		log:           log.NewHelper(log.With(logger, "module", "biz.dashboard")),
		dashboardRepo: dashboardRepo,
		chartRepo:     chartRepo,
	}
}

// CreateChart 创建图表
func (l *DashboardBiz) CreateChart(ctx context.Context, chartInfo *bo.MyChartBO) (*bo.MyChartBO, error) {
	return l.chartRepo.Create(ctx, chartInfo)
}

// UpdateChartById 更新图表
func (l *DashboardBiz) UpdateChartById(ctx context.Context, chartId uint32, chartInfo *bo.MyChartBO) (*bo.MyChartBO, error) {
	return l.chartRepo.Update(ctx, chartInfo, basescopes.InIds(chartId))
}

// DeleteChartById 删除图表
func (l *DashboardBiz) DeleteChartById(ctx context.Context, chartId uint32) error {
	return l.chartRepo.Delete(ctx, basescopes.InIds(chartId))
}

// GetChartDetail 获取图表详情
func (l *DashboardBiz) GetChartDetail(ctx context.Context, chartId uint32) (*bo.MyChartBO, error) {
	return l.chartRepo.Get(ctx, basescopes.InIds(chartId))
}

// ListChartByPage 查询图表列表
func (l *DashboardBiz) ListChartByPage(ctx context.Context, pgInfo basescopes.Pagination, scopes ...basescopes.ScopeMethod) ([]*bo.MyChartBO, error) {
	return l.chartRepo.List(ctx, pgInfo, scopes...)
}

// CreateDashboard 创建dashboard
func (l *DashboardBiz) CreateDashboard(ctx context.Context, dashboardInfo *bo.MyDashboardConfigBO) (*bo.MyDashboardConfigBO, error) {
	return l.dashboardRepo.Create(ctx, dashboardInfo)
}

// UpdateDashboardById 更新dashboard
func (l *DashboardBiz) UpdateDashboardById(ctx context.Context, dashboardId uint32, dashboardInfo *bo.MyDashboardConfigBO) (*bo.MyDashboardConfigBO, error) {
	return l.dashboardRepo.Update(ctx, dashboardInfo, basescopes.InIds(dashboardId))
}

// DeleteDashboardById 删除dashboard
func (l *DashboardBiz) DeleteDashboardById(ctx context.Context, dashboardId uint32) error {
	return l.dashboardRepo.Delete(ctx, basescopes.InIds(dashboardId))
}

// GetDashboardById 获取dashboard详情
func (l *DashboardBiz) GetDashboardById(ctx context.Context, dashboardId uint32) (*bo.MyDashboardConfigBO, error) {
	return l.dashboardRepo.Get(ctx, basescopes.InIds(dashboardId), basescopes.MyDashboardConfigPreloadCharts)
}

// ListDashboard 查询dashboard列表
func (l *DashboardBiz) ListDashboard(ctx context.Context, pgInfo basescopes.Pagination, scopes ...basescopes.ScopeMethod) ([]*bo.MyDashboardConfigBO, error) {
	return l.dashboardRepo.List(ctx, pgInfo, scopes...)
}
