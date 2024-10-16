package userbiz

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type userVerifyForgotPasswordRepo interface {
	SetVerifyForgotPassword(ctx context.Context, data *usermodel.UserSetVerifyForgotPassword) error
}

type userVerifyForgotPasswordBiz struct {
	repo userVerifyForgotPasswordRepo
}

func NewUserVerifyForgotPasswordBiz(repo userVerifyForgotPasswordRepo) *userVerifyForgotPasswordBiz {
	return &userVerifyForgotPasswordBiz{repo: repo}
}

func (b *userVerifyForgotPasswordBiz) SetVerifyForgotPassword(ctx context.Context, data *usermodel.UserSetVerifyForgotPassword) error {
	return b.repo.SetVerifyForgotPassword(ctx, data)
}
