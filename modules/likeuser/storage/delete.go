package likestorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

func (s *sqlStore) DeleteLike(ctx context.Context, userId int, restaurantId int) error {
	db := s.db.Table(userlikerestaurantmodel.Like{}.TableName())
	if err := db.Where("user_id = ? and restaurant_id = ?", userId, restaurantId).Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
