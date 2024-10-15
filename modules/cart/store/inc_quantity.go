package cartstore

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
	"gorm.io/gorm"
)

func (s *sqlStore) IncreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error {
	db := s.db.Table(data.TableName())
	if err := db.Where(data).Update("quantity", gorm.Expr("`quantity`+ ?", 1)).Error; err != nil {
		return common.ErrCannotUpdateEntity("cart", err)
	}
	return nil
}
