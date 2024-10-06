package foodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
	"gorm.io/gorm"
)

func (s *sqlStore) IncLikeFood(context context.Context, foodId int) error {
	db := s.db.Table(foodmodel.EntityName)

	if err := db.Where("id = ?", foodId).Update("like", gorm.Expr("`like` + ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
