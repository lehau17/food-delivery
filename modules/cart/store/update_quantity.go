package cartstore

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateQuantity(ctx context.Context, data *cartmodel.CartCreate) error {
	db := s.db.Table(data.TableName())
	if err := db.Where("user_id = ? and food_id = ?", data.UserId, data.FoodId).Update("quantity", gorm.Expr("`quantity`+ ?", data.Quantity)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
