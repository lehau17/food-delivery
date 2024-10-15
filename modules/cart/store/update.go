package cartstore

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

func (s *sqlStore) UpdateCart(ctx context.Context, newCart *cartmodel.Cart) error {
	db := s.db.Table(newCart.TableName())
	if err := db.Save(newCart).Error; err != nil {
		return common.ErrCannotUpdateEntity("cart", err)
	}
	return nil
}
