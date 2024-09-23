package restaurentstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

func (s *sqlStore) CreateRestaurant(context context.Context, data *restaurentmodel.RestaurantCreate) error {

	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
