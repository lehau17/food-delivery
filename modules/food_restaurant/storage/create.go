package foodrestaurantstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodrestaurantmodel "github.com/lehau17/food_delivery/modules/food_restaurant/model"
)

func (s *sqlStore) CreateRestaurantFood(ctx context.Context, data *foodrestaurantmodel.FoodsRestaurantCreate) error {
	db := s.db.Table(data.TableName())
	if err := db.Create(&data).Error; err != nil {
		return common.ErrCannotCreateEntity("Food restaurant", err)
	}
	return nil
}
