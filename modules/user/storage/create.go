package userstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	// Begin transaction
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}
	//Commit
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}
	return nil
}
