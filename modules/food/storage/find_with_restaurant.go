package foodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) FindFoodWithRestaurant(ctx context.Context, foodId int, userId int) (*foodmodel.Food, error) {
	var food foodmodel.Food

	if err := s.db.Table(foodmodel.EntityName).
		Preload("Restaurant").
		Where("foods.id = ? AND restaurants.user_id = ?", foodId, userId).
		Find(&food).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return &food, nil
}
