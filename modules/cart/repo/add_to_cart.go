package cartrepo

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartAddToStore interface {
	CreateCart(ctx context.Context, data *cartmodel.CartCreate) error
	FindCart(ctx context.Context, conditions interface{}) (*cartmodel.Cart, error)
	UpdateQuantity(ctx context.Context, data *cartmodel.CartCreate) error
}

type CartAddToRepo struct {
	store CartAddToStore
}

func NewCartAddToRepo(store CartAddToStore) *CartAddToRepo {
	return &CartAddToRepo{store: store}
}

func (r *CartAddToRepo) AddToCart(ctx context.Context, data *cartmodel.CartCreate) error {
	// check cart
	foundCart, err := r.store.FindCart(ctx, map[string]interface{}{"food_id": data.FoodId, "user_id": data.UserId})
	if err != nil {
		return err
	}
	// Not found => Create
	if foundCart == nil {
		return r.store.CreateCart(ctx, data)
	} else {
		return r.store.UpdateQuantity(ctx, data)
	}
}
