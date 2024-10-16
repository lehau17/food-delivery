package userstorage

import (
	"context"
	"time"

	"github.com/lehau17/food_delivery/common"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

func (s *sqlStore) ForgotPassword(ctx context.Context, data *usermodel.UserSetVerifyForgotPassword) (string, error) {
	rbd := s.rdb
	err := rbd.Set(ctx, "verify_forgot_password:"+data.Email, data.Token, 300*time.Second).Err()
	if err != nil {
		return "", common.ErrCache(err)
	}
	return "http://locahost:8080/user/verify/forgot-password?Email=" + data.Email + "&Token=" + data.Token, nil
}
