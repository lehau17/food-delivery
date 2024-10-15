package cartrepo

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartDescreaseQuantityStore interface {
	FindCart(ctx context.Context, conditions interface{}) (*cartmodel.Cart, error)
	DescreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error
}

type cartDescreaseQuantityRepo struct {
	store CartDescreaseQuantityStore
}

func NewCartDescreaseQuantityRepo(store CartDescreaseQuantityStore) *cartDescreaseQuantityRepo {
	return &cartDescreaseQuantityRepo{store: store}
}

func (r *cartDescreaseQuantityRepo) DescreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error {
	foundCart, err := r.store.FindCart(ctx, map[string]interface{}{"food_id": data.FoodId, "user_id": data.UserId})
	if err != nil {
		return err
	}
	if foundCart == nil {
		return cartmodel.ErrCartNotFound
	}
	// foundCart.Quantity = data.Quantity
	return r.store.DescreaseQuantity(ctx, data)
}
