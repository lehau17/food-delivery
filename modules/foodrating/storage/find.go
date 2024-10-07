package foodratingstorage

import (
	"context"
	"errors"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

func (s *sqlStore) FindFoodRating(ctx context.Context,
	conditions map[string]interface{},
	morePreload ...string) (*foodratingmodel.FoodRating, error) {

	db := s.db.Table(foodratingmodel.EntityName)
	for i := range morePreload {
		db = db.Preload(morePreload[i])
	}
	var food foodratingmodel.FoodRating
	if err := db.Where(conditions).First(&food).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, common.ErrRecordNotFound(errors.New("food rating not found"))
		}
		return nil, common.ErrDb(err)
	}
	return &food, nil
}
