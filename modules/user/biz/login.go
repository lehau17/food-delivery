package userbiz

import (
	"context"

	"github.com/lehau17/food_delivery/components/jwtprovider"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type UserLoginStore interface {
	Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type userLoginBiz struct {
	jwtprovider jwtprovider.TokenProvider
	store       UserLoginStore
	hash        Hasher
	expiry      int
}

func NewUserLoginBiz(store UserLoginStore, hash Hasher, provider jwtprovider.TokenProvider, expiry int) *userLoginBiz {
	return &userLoginBiz{store: store, hash: hash, jwtprovider: provider, expiry: expiry}
}

func (b *userLoginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*jwtprovider.Token, error) {
	foundUser, err := b.store.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, err
	}
	if foundUser == nil {
		return nil, usermodel.ErrUserNotFound
	}
	// get salt and comparing against
	salt := foundUser.Salt
	hashPassword := b.hash.Hash(data.Password + salt)
	if hashPassword != foundUser.Password {
		return nil, usermodel.ErrUserLoginFail
	}

	// generate JWT token
	token, err := b.jwtprovider.SignToken(&jwtprovider.TokenPayload{Uid: foundUser.Id, Role: foundUser.Role}, b.expiry)
	if err != nil {
		return nil, jwtprovider.ErrCreateToken
	}
	return token, nil
}
