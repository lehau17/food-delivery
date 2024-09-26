package likestorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

func (s *sqlStore) FindLike(context context.Context, userId int, resId int) (*userlikerestaurantmodel.Like, error) {
	db := s.db.Table(userlikerestaurantmodel.Like{}.TableName())

	var likeuser userlikerestaurantmodel.Like
	if err := db.Where("user_id = ? and restaurant_id = ?", userId, resId).First(&likeuser).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return &likeuser, nil
}
