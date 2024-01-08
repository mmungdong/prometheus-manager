package chatgroup

import (
	"context"

	query "github.com/aide-cloud/gorm-normalize"
	"github.com/go-kratos/kratos/v2/log"

	"prometheus-manager/api/perrors"
	"prometheus-manager/app/prom_server/internal/biz/bo"
	"prometheus-manager/app/prom_server/internal/biz/do"
	"prometheus-manager/app/prom_server/internal/biz/do/basescopes"
	"prometheus-manager/app/prom_server/internal/biz/repository"
	"prometheus-manager/app/prom_server/internal/data"
	"prometheus-manager/pkg/util/slices"
)

var _ repository.ChatGroupRepo = (*chatGroupRepoImpl)(nil)

var (
	// ErrNoCondition 不允许无条件操作DB
	ErrNoCondition = perrors.ErrorUnknown("no condition")
)

type chatGroupRepoImpl struct {
	repository.UnimplementedChatGroupRepo
	log  *log.Helper
	data *data.Data
}

func (l *chatGroupRepoImpl) Create(ctx context.Context, chatGroup *bo.ChatGroupBO) (*bo.ChatGroupBO, error) {
	newChatGroup := chatGroup.ToModel()
	if err := l.data.DB().WithContext(ctx).Create(newChatGroup).Error; err != nil {
		return nil, err
	}
	return bo.ChatGroupModelToBO(newChatGroup), nil
}

func (l *chatGroupRepoImpl) Get(ctx context.Context, scopes ...query.ScopeMethod) (*bo.ChatGroupBO, error) {
	var chatGroupDetail do.PromAlarmChatGroup
	if err := l.data.DB().WithContext(ctx).Scopes(scopes...).First(&chatGroupDetail).Error; err != nil {
		return nil, err
	}

	return bo.ChatGroupModelToBO(&chatGroupDetail), nil
}

func (l *chatGroupRepoImpl) Find(ctx context.Context, scopes ...query.ScopeMethod) ([]*bo.ChatGroupBO, error) {
	var chatGroupList []*do.PromAlarmChatGroup
	if err := l.data.DB().WithContext(ctx).Scopes(scopes...).Find(&chatGroupList).Error; err != nil {
		return nil, err
	}
	list := slices.To(chatGroupList, func(i *do.PromAlarmChatGroup) *bo.ChatGroupBO {
		return bo.ChatGroupModelToBO(i)
	})
	return list, nil
}

func (l *chatGroupRepoImpl) Update(ctx context.Context, chatGroup *bo.ChatGroupBO, scopes ...query.ScopeMethod) error {
	if len(scopes) == 0 {
		return ErrNoCondition
	}
	return l.data.DB().WithContext(ctx).Scopes(scopes...).Updates(chatGroup.ToModel()).Error
}

func (l *chatGroupRepoImpl) Delete(ctx context.Context, scopes ...query.ScopeMethod) error {
	if len(scopes) == 0 {
		return ErrNoCondition
	}
	return l.data.DB().WithContext(ctx).Scopes(scopes...).Delete(&do.PromAlarmChatGroup{}).Error
}

func (l *chatGroupRepoImpl) List(ctx context.Context, pgInfo query.Pagination, scopes ...query.ScopeMethod) ([]*bo.ChatGroupBO, error) {
	var chatGroupList []*do.PromAlarmChatGroup
	if err := l.data.DB().WithContext(ctx).Scopes(append(scopes, basescopes.Page(pgInfo))...).Find(&chatGroupList).Error; err != nil {
		return nil, err
	}
	if pgInfo != nil {
		var total int64
		if err := l.data.DB().WithContext(ctx).Scopes(scopes...).Model(&do.PromAlarmChatGroup{}).Count(&total).Error; err != nil {
			return nil, err
		}
		pgInfo.SetTotal(total)
	}

	return slices.To(chatGroupList, func(i *do.PromAlarmChatGroup) *bo.ChatGroupBO {
		return bo.ChatGroupModelToBO(i)
	}), nil
}

func NewChatGroupRepo(d *data.Data, logger log.Logger) repository.ChatGroupRepo {
	return &chatGroupRepoImpl{
		log:  log.NewHelper(log.With(logger, "module", "data.repository.chat_group")),
		data: d,
	}
}
