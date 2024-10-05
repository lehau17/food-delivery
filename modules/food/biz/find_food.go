package foodbiz

import (
	"context"

	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodFindRepo interface {
	FindFood(ctx context.Context, id int, filter *foodmodel.Filter, morePreload ...string) (*foodmodel.Food, error)
}

type FoodFindBiz struct {
	repo FoodFindRepo
}

func NewFoodFindBiz(repo FoodFindRepo) *FoodFindBiz {
	return &FoodFindBiz{repo: repo}
}

func (b *FoodFindBiz) FindFood(ctx context.Context, id int, filter *foodmodel.Filter) (*foodmodel.Food, error) {
	return b.repo.FindFood(ctx, id, filter, "Restaurant", "Category")
}
