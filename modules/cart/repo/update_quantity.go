package cartrepo

import (
	"context"
	"errors"

	"github.com/lehau17/food_delivery/common"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartUpdateQuantityStore interface {
	FindCart(ctx context.Context, conditions interface{}) (*cartmodel.Cart, error)
	UpdateCart(ctx context.Context, newCart *cartmodel.Cart) error
}

type cartUpdateQuantityRepo struct {
	store CartUpdateQuantityStore
}

func NewCartUpdateQuantityRepo(store CartUpdateQuantityStore) *cartUpdateQuantityRepo {
	return &cartUpdateQuantityRepo{store: store}
}

func (r *cartUpdateQuantityRepo) UpdateCart(ctx context.Context, data *cartmodel.CartUpdate) error {
	foundCart, err := r.store.FindCart(ctx, map[string]interface{}{"food_id": data.FoodId, "user_id": data.UserId})
	if err != nil {
		return err
	}
	if foundCart == nil {
		return cartmodel.ErrCartNotFound
	}
	if foundCart.Quantity < data.Quantity {
		return common.ErrInvalidRequest(errors.New("invalid request"))
	}
	foundCart.Quantity = data.Quantity
	return r.store.UpdateCart(ctx, foundCart)
}
