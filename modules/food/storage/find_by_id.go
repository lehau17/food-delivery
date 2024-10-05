package foodstorage

import (
	"context"

	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) FindFoodById(ctx context.Context, id int, filter *foodmodel.Filter, morePreload ...string) (*foodmodel.Food, error) {
	db := s.db.Table(foodmodel.EntityName)
	if len(morePreload) > 0 {
		for i := range morePreload {
			db = db.Preload(morePreload[i])
		}
	}
	var food foodmodel.Food
	if err := db.Where("id = ? and status = ?", id, *filter.Status).First(&food).Error; err != nil {
		return nil, err
	}

	return &food, nil
}
