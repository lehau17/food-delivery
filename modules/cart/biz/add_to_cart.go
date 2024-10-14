package cartbiz

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartAddToRepo interface {
	AddToCart(ctx context.Context, data *cartmodel.CartCreate) error
}

type cartAddToBiz struct {
	repo CartAddToRepo
}

func NewCartAddToBiz(repo CartAddToRepo) *cartAddToBiz {
	return &cartAddToBiz{repo: repo}
}

func (b *cartAddToBiz) AddToCart(ctx context.Context, data *cartmodel.CartCreate) error {
	return b.repo.AddToCart(ctx, data)
}
