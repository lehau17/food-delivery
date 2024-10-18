package orderstore

import (
	"context"

	ordermodel "github.com/lehau17/food_delivery/modules/order/model"
)

func (s *sqlStore) FindOrder(ctx context.Context, conditions interface{}, moreField ...string) (*ordermodel.Order, error) {
	db := s.db.Table(ordermodel.EntityName)
	for i := range moreField {
		db = db.Preload(moreField[i])
	}
	var orderFound ordermodel.Order
	if err := db.Where(conditions).First(&orderFound).Error; err != nil {
		if err.Error() == "record not found" {
		}
	}
	return nil, nil
}
