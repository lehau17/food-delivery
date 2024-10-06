package likefoodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

func (s *sqlStore) GetListFoodUserLiked(ctx context.Context, userId int) ([]int, error) {
	db := s.db.Table(likefoodmodel.EntityName)
	var ids []int
	if err := db.Where("user_id = ?", userId).Select("food_id").Find(&ids).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return ids, nil
}
