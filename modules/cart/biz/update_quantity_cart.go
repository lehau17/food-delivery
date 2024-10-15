package cartbiz

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartUpdateQuantityRepo interface {
	UpdateCart(ctx context.Context, data *cartmodel.CartUpdate) error
}

type cartUpdateQuantityBiz struct {
	repo CartUpdateQuantityRepo
}

func NewCartUpdateQuantityBiz(repo CartUpdateQuantityRepo) *cartUpdateQuantityBiz {
	return &cartUpdateQuantityBiz{repo: repo}
}

func (b *cartUpdateQuantityBiz) UpdateQuantity(ctx context.Context, data *cartmodel.CartUpdate) error {
	return b.repo.UpdateCart(ctx, data)
}
