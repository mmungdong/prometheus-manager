package strategyGroup

import (
	"errors"
	"os"

	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"prometheus-manager/pkg/strategy"
)

type (
	ExportReq struct {
		ID uint `form:"id" binding:"required"`
	}
)

// GetExport ...
func (l *StrategyGroup) GetExport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ExportReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ginplus.Logger().Error("bind query error", zap.Error(err))
			ctx.AbortWithStatus(400)
			return
		}

		groupDetail, err := l.groupRepository.GetStrategyGroupDetailById(ctx, req.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.AbortWithStatus(404)
				return
			}
			ctx.AbortWithStatus(500)
			return
		}

		rules := make([]*strategy.Rule, 0, len(groupDetail.PromStrategies))
		for _, promStrategy := range groupDetail.PromStrategies {
			rules = append(rules, &strategy.Rule{
				Alert:       promStrategy.Alert,
				Expr:        promStrategy.Expr,
				For:         promStrategy.For,
				Labels:      promStrategy.Labels,
				Annotations: promStrategy.Annotations,
			})
		}

		// 把groupDetail转换成GroupFile
		resp, err := l.groupRepository.ExportStrategyGroupFile(ctx, &strategy.Group{
			Name:  groupDetail.Name,
			Rules: rules,
		})
		if err != nil {
			ginplus.Logger().Error("export strategy group file error", zap.Error(err))
			ctx.AbortWithStatus(500)
			return
		}
		defer func() {
			_ = os.Remove(resp.Name())
		}()

		// 下载文件
		ctx.Writer.Header().Add("Content-Disposition", "attachment; filename="+groupDetail.Name+".yaml")
		ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
		ctx.File(resp.Name())
	}
}
