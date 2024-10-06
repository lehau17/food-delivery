package foodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) GetFoods(ctx context.Context,
	conditions map[string]interface{},
	filter *foodmodel.Filter,
	paging *common.PagingCursor,
	moreLoad ...string) ([]foodmodel.Food, error) {
	db := s.db.Table(foodmodel.EntityName)
	for i := range moreLoad {
		db = db.Preload(moreLoad[i])
	}
	if filter.Status != nil {
		db = db.Where("status = ?", *filter.Status)
	}
	if paging.Cursor != "" {
		db = db.Where("id > ?", paging.RealCursor)
	}
	var data []foodmodel.Food
	if err := db.Where(conditions).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return data, nil
}
