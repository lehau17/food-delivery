package restaurantratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

func (s *sqlStore) DeleteRestaurantRating(ctx context.Context, id int) error {
	db := s.db.Table(restaurantratingmodel.EntityName)
	if err := db.Where("id = ? and status = 1", id).Update("status", 0).Error; err != nil {
		if err.Error() == "record not found" {
			return restaurantratingmodel.ErrRestaurantRatingNotExist
		}
		return common.ErrCannotDeleteEntity("Food rating", err)
	}
	return nil
}
