package userrepo

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type userChangePasswordStore interface {
	Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	ChangePassword(ctx context.Context, data *usermodel.ChangePassword) error
}

type hash interface {
	Hash(data string) string
}

type userChangePasswordRepo struct {
	store userChangePasswordStore
	hash  hash
}

func NewUserChangePasswordRepo(store userChangePasswordStore, hash hash) *userChangePasswordRepo {
	return &userChangePasswordRepo{store: store, hash: hash}
}

func (r *userChangePasswordRepo) ChangePassword(ctx context.Context, data *usermodel.ChangePassword) error {
	foundUser, err := r.store.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return err
	}
	if foundUser == nil {
		return usermodel.ErrUserNotExists
	}
	if foundUser.Status == 0 {
		return usermodel.ErrUserDisable
	}
	// check hash password
	newHash := r.hash.Hash(data.Password + foundUser.Salt)
	if newHash == foundUser.Password {
		return usermodel.ErrNewPassWordIsNotBangOldPassword
	}
	data.HashPassword = newHash
	return r.store.ChangePassword(ctx, data)
}
