package cartstore

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

func (s *sqlStore) FindCart(ctx context.Context, conditions interface{}) (*cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.EntityName)
	var data cartmodel.Cart
	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, common.ErrDb(err)
	}
	return &data, nil
}
