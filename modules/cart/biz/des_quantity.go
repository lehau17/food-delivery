package cartbiz

import (
	"context"

	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
)

type CartDescreaseQuantityRepo interface {
	DescreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error
}

type cartDescreaseQuantityBiz struct {
	repo CartDescreaseQuantityRepo
}

func NewCartDescreaseQuantityBiz(repo CartDescreaseQuantityRepo) *cartDescreaseQuantityBiz {
	return &cartDescreaseQuantityBiz{repo: repo}
}

func (b *cartDescreaseQuantityBiz) DescreaseQuantity(ctx context.Context, data *cartmodel.CartChangeQuantity) error {
	return b.repo.DescreaseQuantity(ctx, data)
}
