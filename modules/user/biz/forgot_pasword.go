package userbiz

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type UserForgotPasswordRepo interface {
	ForgotPassword(ctx context.Context, data *usermodel.UserForgotPassword) error
}

type userForgotPasswordBiz struct {
	repo UserForgotPasswordRepo
}

func NewUserForgotPasswordBiz(repo UserForgotPasswordRepo) *userForgotPasswordBiz {
	return &userForgotPasswordBiz{repo: repo}
}

func (b *userForgotPasswordBiz) ForgotPassword(ctx context.Context, data *usermodel.UserForgotPassword) error {
	return b.repo.ForgotPassword(ctx, data)
}
