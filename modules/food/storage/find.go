package foodstorage

import (
	"context"

	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) FindFood(ctx context.Context, coditions map[string]interface{}, morePreload ...string) ([]foodmodel.Food, error) {
	db := s.db.Table(foodmodel.EntityName)
	if len(morePreload) > 0 {
		for i := range morePreload {
			db = db.Preload(morePreload[i])
		}
	}
	var foods []foodmodel.Food
	if err := db.Where(coditions).Find(&foods).Error; err != nil {
		return nil, err
	}

	return foods, nil
}
