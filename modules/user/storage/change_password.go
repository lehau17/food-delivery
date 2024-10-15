package userstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

func (s *sqlStore) ChangePassword(ctx context.Context, data *usermodel.ChangePassword) error {
	db := s.db.Table(data.TableName())
	if err := db.Where("email = ?", data.Email).Update("password = ?", data.HashPassword).Error; err != nil {
		return common.ErrCannotUpdateEntity("users", err)
	}
	return nil
}
