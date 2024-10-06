package likefoodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

func (s *sqlStore) DeleteLike(ctx context.Context, foodId int, userId int) error {
	db := s.db.Table(likefoodmodel.EntityName)
	if err := db.Where("food_id = ? and user_id = ?").Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
