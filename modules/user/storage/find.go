package userstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var user usermodel.User
	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(err)
		}
		return nil, common.ErrDb(err)
	}
	return &user, nil
}
