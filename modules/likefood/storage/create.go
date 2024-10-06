package likefoodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

func (s *sqlStore) CreateLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error {
	db := s.db.Table(data.TableName())
	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
