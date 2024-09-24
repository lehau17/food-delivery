package userbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type UserRegisterStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type Hash interface {
	Hash(data string) string
}

type userRegisterBiz struct {
	UserRegisterStore UserRegisterStore
	Hash              Hash
}

func NewUserRegisterBiz(store UserRegisterStore, hash Hash) *userRegisterBiz {
	return &userRegisterBiz{UserRegisterStore: store, Hash: hash}
}

func (b *userRegisterBiz) RegisterUser(ctx context.Context, data *usermodel.UserCreate) error {
	//check user exists
	foundUser, err := b.UserRegisterStore.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		if error, ok := err.(*common.AppError); ok {
			if error.Key != "ErrRecordNotFound" {
				return err
			}
		}
	}
	if foundUser != nil {
		if foundUser.Status == 0 {
			return usermodel.ErrUserDisable
		}
		return usermodel.ErrUserExists
	}

	//create user
	salt := common.GetSalt(50)
	data.Salt = salt
	hashPassword := b.Hash.Hash(data.Password + salt)
	data.Password = hashPassword
	data.Role = "user"
	data.Status = 1
	if err := b.UserRegisterStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("users", err)
	}
	return nil

}
