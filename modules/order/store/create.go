package orderstore

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	ordermodel "github.com/lehau17/food_delivery/modules/order/model"
)

func (s *sqlStore) CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) (int, error) {
	db := s.db.Table(data.TableName())
	if err := db.Create(data).Error; err != nil {
		return 0, common.ErrCannotCreateEntity("Order", err)
	}
	return data.Id, nil
}
