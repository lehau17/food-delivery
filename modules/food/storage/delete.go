package foodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) DeleteFood(ctx context.Context, id int) error {
	db := s.db.Table(foodmodel.EntityName)
	if err := db.Where("id = ? and status = ?", id, 1).Update("status", 0).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
