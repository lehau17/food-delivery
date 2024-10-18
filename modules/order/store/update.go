package orderstore

import (
	"context"

	ordermodel "github.com/lehau17/food_delivery/modules/order/model"
)

func (s *sqlStore) UpdateOrder(ctx context.Context, data *ordermodel.Order) error {
	db := s.db.Table(data.TableName())
	if err := db.Save(data).Error; err != nil {
		return err
	}
	return nil
}
