package restaurantratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

func (s *sqlStore) CreateRestaurantRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error {
	db := s.db.Table(data.TableName())
	if err := db.Create(data).Error; err != nil {
		return common.ErrCannotCreateEntity("Food rating", err)
	}
	return nil
}
