package foodbiz

import (
	"context"

	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodUpdateRepo interface {
	UpdateFood(ctx context.Context, data *foodmodel.FoodUpdate, id int) error
}

type FoodUpdateBiz struct {
	repo FoodUpdateRepo
}

func NewFoodUpdateBiz(repo FoodUpdateRepo) *FoodUpdateBiz {
	return &FoodUpdateBiz{repo: repo}
}

func (b *FoodUpdateBiz) UpdateFood(ctx context.Context, data *foodmodel.FoodUpdate, id int) error {
	return b.repo.UpdateFood(ctx, data, id)
}
