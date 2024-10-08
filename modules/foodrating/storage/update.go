package foodratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

func (s *sqlStore) UpdateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingUpdate, oldData *foodratingmodel.FoodRating) error {
	db := s.db.Table(data.TableName())
	if data.Comment != "" {
		oldData.Comment = data.Comment
	}

	if data.Point > 0 {
		oldData.Point = data.Point
	}
	if err := db.Save(oldData).Error; err != nil {
		return common.ErrCannotUpdateEntity(data.TableName(), err)
	}
	return nil
}
