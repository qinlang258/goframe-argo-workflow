package logic

import (
	"context"
	"goframe-argo-workflow/internal/dao"
	"goframe-argo-workflow/internal/model/entity"
)

type sUser struct {
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) UserCreate(ctx context.Context, userinfo *entity.User) error {
	_, err := dao.User.Ctx(ctx).Save(userinfo)
	if err != nil {
		return err
	}

	return err
}
