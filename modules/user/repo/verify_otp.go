package userrepo

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type UserVerifyStore interface {
	Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	EnableUser(ctx context.Context, data *usermodel.UserVerifyOtp) error
}

type UserVerifyBiz struct {
	store UserVerifyStore
}

func NewUserVerifyBiz(store UserVerifyStore) *UserVerifyBiz {
	return &UserVerifyBiz{store: store}
}

func (b *UserVerifyBiz) VerifyOtp(ctx context.Context, data *usermodel.UserVerifyOtp) error {
	// check Email
	foundUser, err := b.store.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return err
	}
	if foundUser == nil {
		return usermodel.ErrUserNotExists
	}
	if foundUser.Status == 0 {
		return usermodel.ErrUserDisable
	}
	if foundUser.Status == 1 {
		return usermodel.ErrUserEnable
	}
	return b.store.EnableUser(ctx, data)
}
