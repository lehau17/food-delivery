package userstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

func (s *sqlStore) GetListUser(ctx context.Context, condition map[string]interface{}) ([]common.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())
	var users []common.User
	if ids, ok := condition["id"]; ok {
		db = db.Where("id in (?)", ids)
	}
	if err := db.Where(condition).Find(&users).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return users, nil
}
