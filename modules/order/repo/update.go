package orderrepo

import (
	"context"

	ordermodel "github.com/lehau17/food_delivery/modules/order/model"
)

type orderUpdateStore interface {
	FindOrder(ctx context.Context, conditions interface{}, moreField ...string) (*ordermodel.Order, error)
	UpdateOrder(ctx context.Context, data *ordermodel.Order) error
}

type orderUpdateRepo struct {
	store orderUpdateStore
}

func NewOrderUpdateRepo(store orderUpdateStore) *orderUpdateRepo {
	return &orderUpdateRepo{store: store}
}

func (r *orderUpdateRepo) UpdateOrder(ctx context.Context, data *ordermodel.Order) error {
	foundOrder, err := r.store.FindOrder(ctx, map[string]interface{}{"id": data.Id})
	if err != nil {
		return err
	}

	return r.store.UpdateOrder(ctx, foundOrder)
}
