package restaurantratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

func (s *sqlStore) FindRestaurantRating(ctx context.Context, conditions map[string]interface{}, moreField ...string) (*restaurantratingmodel.RestaurantRating, error) {
	db := s.db.Table(restaurantratingmodel.EntityName)
	var data restaurantratingmodel.RestaurantRating
	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, restaurantratingmodel.ErrRestaurantRatingNotExist
		}
		return nil, common.ErrDb(err)
	}
	return &data, nil
}
