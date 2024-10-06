package likefoodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

func (s *sqlStore) GetListUserLikeFood(ctx context.Context, foodId int) ([]int, error) {
	db := s.db.Table(likefoodmodel.EntityName)
	var arrIds []int
	if err := db.Where("food_id = ?", foodId).Select("user_id").Find(&arrIds).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return arrIds, nil

}
