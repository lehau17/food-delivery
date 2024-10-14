package userstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

func (s *sqlStore) EnableUser(ctx context.Context, data *usermodel.UserVerifyOtp) error {
	db := s.db.Table(usermodel.EntityName)
	//check OTP
	otp := s.rdb.Get(ctx, "otp:"+data.Email)
	if otp.Val() != data.Otp {
		return usermodel.ErrOtp
	}
	if err := db.Find("email = ?", data.Email).Update("status", 1).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
