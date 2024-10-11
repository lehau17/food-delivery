package foodrestaurantstorage

import (
	"github.com/lehau17/food_delivery/common"
	foodrestaurantmodel "github.com/lehau17/food_delivery/modules/food_restaurant/model"
)

func (s *sqlStore) DeleteRestaurantFood(ctx, resId int, foodId int) error {
	db := s.db.Table(foodrestaurantmodel.EntityName)
	if err := db.Where("restaurant_id = ? and food_id = ? and status = 1", resId, foodId).Update("status", 0).Error; err != nil {
		if err.Error() == "record not found" {
			return foodrestaurantmodel.ErrFoodRatingNotExist
		}
		return common.ErrCannotDeleteEntity("Food restaurant", err)
	}
	return nil
}
