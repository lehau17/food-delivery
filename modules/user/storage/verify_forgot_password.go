package userstorage

import (
	"context"

	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

func (s *sqlStore) SetVerifyForgotPassword(ctx context.Context, data *usermodel.UserSetVerifyForgotPassword) error {
	rbd := s.rdb
	stringCmd := rbd.Get(ctx, "verify_forgot_password:"+data.Email)
	if stringCmd.Val() == "" {
		return usermodel.ErrUserDontCanChangePassword
	} else if stringCmd.Val() != data.Token {
		return usermodel.ErrUserDontCanChangePassword
	}
	return nil
}
