package orderrepo

import (
	"context"

	ordermodel "github.com/lehau17/food_delivery/modules/order/model"
)

type orderCreateStore interface {
	CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) (int, error)
}

type orderCreateRepo struct {
	store orderCreateStore
}

func NewOrderCreateRepo(store orderCreateStore) *orderCreateRepo {
	return &orderCreateRepo{store: store}
}

func (r *orderCreateRepo) CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) (int, error) {
	return r.store.CreateOrder(ctx, data)
}
