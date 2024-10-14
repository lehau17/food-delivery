package cartstore

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

func (s *sqlStore) CreateCart(ctx context.Context, data *cartmodel.CartCreate) error {
	db := s.db.Table(data.TableName())
	if err := db.Create(&data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
