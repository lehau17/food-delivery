package userrepo

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type userVerifyForgotPasswordStore interface {
	SetVerifyForgotPassword(ctx context.Context, data *usermodel.UserSetVerifyForgotPassword) error
}

type userVerifyForgotPasswordRepo struct {
	repo userVerifyForgotPasswordStore
}

func NewUserVerifyForgotPasswordRepo(store userVerifyForgotPasswordStore) *userVerifyForgotPasswordRepo {
	return &userVerifyForgotPasswordRepo{repo: store}
}

func (r *userVerifyForgotPasswordRepo) SetVerifyForgotPassword(ctx context.Context, data *usermodel.UserSetVerifyForgotPassword) error {
	return r.repo.SetVerifyForgotPassword(ctx, data)
}
