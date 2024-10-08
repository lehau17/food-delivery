package foodratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

func (s *sqlStore) GetListFoodRating(
	ctx context.Context,
	conditions map[string]interface{},
	filter *foodratingmodel.Filter,
	paging *common.PagingCursor,
	morePreload ...string) ([]foodratingmodel.FoodRating, error) {
	db := s.db.Table(foodratingmodel.EntityName)
	for i := range morePreload {
		db = db.Preload(morePreload[i])
	}
	if paging.RealCursor != 0 {
		db = db.Where("id < ?", paging.RealCursor)
	}
	var foods []foodratingmodel.FoodRating
	if err := db.Where(filter).Where(conditions).Limit(paging.Limit).Find(&foods).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return foods, nil
}
