package foodrepo

import (
	"context"

	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodFindStore interface {
	FindFoodById(ctx context.Context, id int, filter *foodmodel.Filter, morePreload ...string) (*foodmodel.Food, error)
}

type FoodFindRepo struct {
	Repo FoodFindStore
}

func NewFoodFindRepo(repo FoodFindStore) *FoodFindRepo {
	return &FoodFindRepo{Repo: repo}
}
func (r *FoodFindRepo) FindFood(ctx context.Context, id int, filter *foodmodel.Filter, morePreload ...string) (*foodmodel.Food, error) {
	return r.Repo.FindFoodById(ctx, id, filter, morePreload...)
}
