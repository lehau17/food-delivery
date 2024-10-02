package restaurentstorage

import (
	"context"
	"log"

	"github.com/lehau17/food_delivery/common"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
	"gorm.io/gorm"
)

func (s *sqlStore) LikeRestaurant(context context.Context, resId int) error {
	db := s.db.Table(restaurentmodel.Restaurant{}.TableName())
	log.Println("resId :>>>>", resId)
	if err := db.Where("id = ?", resId).Update("like", gorm.Expr("`like` + ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
