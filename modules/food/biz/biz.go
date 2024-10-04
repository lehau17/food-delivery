package foodbiz

import (
	"context"

	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodCreateRepo interface {
	CreateFood(ctx context.Context, data *foodmodel.FoodCreate, userId int) error
}

type FoodCreateBiz struct {
	repo FoodCreateRepo
}

func NewFoodCreateBiz(repo FoodCreateRepo) *FoodCreateBiz {
	return &FoodCreateBiz{repo: repo}
}

func (b *FoodCreateBiz) CreateFood(ctx context.Context, data *foodmodel.FoodCreate, userId int) error {
	return b.repo.CreateFood(ctx, data, userId)
}
