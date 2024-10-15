package cartrepo

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartIncreaseQuantityStore interface {
	FindCart(ctx context.Context, conditions interface{}) (*cartmodel.Cart, error)
	IncreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error
}

type cartIncreaseQuantityRepo struct {
	store CartIncreaseQuantityStore
}

func NewCartIncreaseQuantityRepo(store CartIncreaseQuantityStore) *cartIncreaseQuantityRepo {
	return &cartIncreaseQuantityRepo{store: store}
}

func (r *cartIncreaseQuantityRepo) IncreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error {
	foundCart, err := r.store.FindCart(ctx, map[string]interface{}{"food_id": data.FoodId, "user_id": data.UserId})
	if err != nil {
		return err
	}
	if foundCart == nil {
		return cartmodel.ErrCartNotFound
	}
	// foundCart.Quantity = data.Quantity
	return r.store.IncreaseQuantity(ctx, data)
}
