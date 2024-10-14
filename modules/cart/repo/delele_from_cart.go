package cartrepo

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartDeleteFromStore interface {
	FindCart(ctx context.Context, conditions interface{}) (*cartmodel.Cart, error)
	DeleteCart(ctx context.Context, data *cartmodel.CartDelete) error
}

type cartDeleteFromRepo struct {
	repo CartDeleteFromStore
}

func NewCartDeleteFromRepo(repo CartDeleteFromStore) *cartDeleteFromRepo {
	return &cartDeleteFromRepo{repo: repo}
}

func (r *cartDeleteFromRepo) DeleteFromCart(ctx context.Context, data *cartmodel.CartDelete) error {
	foundCart, err := r.repo.FindCart(ctx, data)
	if err != nil {
		return err
	}
	if foundCart == nil {
		return cartmodel.ErrCartNotFound
	}
	return r.repo.DeleteCart(ctx, data)
}
