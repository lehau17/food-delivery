package foodratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

func (s *sqlStore) CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error {
	db := s.db.Table(data.TableName())
	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
