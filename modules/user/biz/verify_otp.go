package userbiz

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type UserVerifyRepo interface {
	VerifyOtp(ctx context.Context, data *usermodel.UserVerifyOtp) error
}

type userVerifyBiz struct {
	repo UserVerifyRepo
}

func NewUserVerifyBiz(repo UserVerifyRepo) *userVerifyBiz {
	return &userVerifyBiz{repo: repo}
}

func (b *userVerifyBiz) VerifyUser(ctx context.Context, data *usermodel.UserVerifyOtp) error {
	return b.repo.VerifyOtp(ctx, data)
}
