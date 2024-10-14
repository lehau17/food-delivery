package cartbiz

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartDeleteFromRepo interface {
	DeleteFromCart(ctx context.Context, data *cartmodel.CartDelete) error
}

type cartDeleteFromBiz struct {
	repo CartDeleteFromRepo
}

func NewCartDeleteFromBiz(repo CartDeleteFromRepo) *cartDeleteFromBiz {
	return &cartDeleteFromBiz{repo: repo}
}

func (b *cartDeleteFromBiz) DeleteFromCart(ctx context.Context, data *cartmodel.CartDelete) error {
	return b.repo.DeleteFromCart(ctx, data)
}
