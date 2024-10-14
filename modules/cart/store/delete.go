package cartstore

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

func (s *sqlStore) DeleteCart(ctx context.Context, data *cartmodel.CartDelete) error {
	db := s.db.Table(data.TableName())
	if err := db.Where(data).Delete(nil).Error; err != nil {
		if err.Error() == "record not found" {
			return cartmodel.ErrCartNotFound
		}
		return common.ErrCannotDeleteEntity("cart", err)
	}
	return nil
}
