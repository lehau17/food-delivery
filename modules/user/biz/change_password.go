package userbiz

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type userChangePasswordRepo interface {
	ChangePassword(ctx context.Context, data *usermodel.ChangePassword) error
}

type userChangePasswordBiz struct {
	repo userChangePasswordRepo
}

func NewUserChangePasswordBiz(repo userChangePasswordRepo) *userChangePasswordBiz {
	return &userChangePasswordBiz{repo: repo}
}

func (r *userChangePasswordBiz) ChangePassword(ctx context.Context, data *usermodel.ChangePassword) error {
	return r.repo.ChangePassword(ctx, data)
}
