package restaurentstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

func (s *sqlStore) FindRestaurant(ctx context.Context, conditions map[string]interface{}) (*restaurentmodel.Restaurant, error) {
	db := s.db.Table(restaurentmodel.Restaurant{}.TableName())
	var data restaurentmodel.Restaurant
	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return &data, nil

}
