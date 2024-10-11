package foodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) CreateFood(ctx context.Context, data *foodmodel.FoodCreate) (int, error) {
	db := s.db.Table(data.TableName())
	if err := db.Create(data).Error; err != nil {
		return 0, common.ErrDb(err)
	}
	return data.ID, nil
}
