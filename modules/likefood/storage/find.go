package likefoodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

func (s *sqlStore) FindLike(ctx context.Context, foodId int, userId int) (*likefoodmodel.LikeFood, error) {
	db := s.db.Table(likefoodmodel.EntityName)
	var foundFood likefoodmodel.LikeFood
	if err := db.Where("food_id = ? and user_id = ?", foodId, userId).First(&foundFood).Error; err != nil {
		if err.Error() == "record not found" {

			return nil, common.ErrRecordNotFound(err)
		}
		return nil, common.ErrDb(err)

	}
	return &foundFood, nil
}
