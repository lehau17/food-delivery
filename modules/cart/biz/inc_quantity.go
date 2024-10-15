package cartbiz

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartIncreaseQuantityRepo interface {
	IncreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error
}

type cartIncreaseQuantityBiz struct {
	repo CartIncreaseQuantityRepo
}

func NewCartIncreaseQuantityBiz(repo CartIncreaseQuantityRepo) *cartIncreaseQuantityBiz {
	return &cartIncreaseQuantityBiz{repo: repo}
}

func (b *cartIncreaseQuantityBiz) IncreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error {
	return b.repo.IncreaseQuantity(ctx, data)
}
