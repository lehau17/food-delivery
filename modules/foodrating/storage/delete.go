package foodratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

func (s *sqlStore) DeleteFoodRating(ctx context.Context, id int, userId int) error {
	db := s.db.Table(foodratingmodel.EntityName)
	if err := db.Where("id = ? and status = ? and user_id = ?", id, 1, userId).Update("status", 0).Error; err != nil {
		if err.Error() == "record not found" {
			return foodratingmodel.ErrFoodRatingNotExists
		}
		return common.ErrDb(err)
	}
	return nil
}
