package restaurantratingstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

func (s *sqlStore) GetListRestaurantRating(ctx context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter, paging *common.PagingCursor, moreField ...string) ([]restaurantratingmodel.RestaurantRating, error) {
	db := s.db.Table(restaurantratingmodel.EntityName)
	var data []restaurantratingmodel.RestaurantRating
	for i := range moreField {
		db = db.Preload(moreField[i])
	}
	if paging.RealCursor != 0 {
		db = db.Where("id < ?", paging.RealCursor)
	}
	if err := db.Where(conditions).Where(filter).Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return data, nil
}
