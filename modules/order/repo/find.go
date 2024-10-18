package orderrepo

import (
	"context"

	ordermodel "github.com/lehau17/food_delivery/modules/order/model"
)

type orderFindStore interface {
	FindOrder(ctx context.Context, conditions interface{}, moreField ...string) (*ordermodel.Order, error)
}

type orderFindRepo struct {
	store orderFindStore
}

func NewOrderFindRepo(store orderFindStore) *orderFindRepo {
	return &orderFindRepo{store: store}
}

// func (r *orderFindRepo) FindOrder(ctx context.Context, data *ordermodel.Order) error {
// 	foundOrder, err := r.store.FindOrder(ctx, map[string]interface{}{"id": data.Id})
// 	if err != nil {
// 		return err
// 	}

// 	return r.store.FindOrder(ctx, foundOrder)
// }
