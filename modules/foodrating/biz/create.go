package foodratingbiz

import (
	"context"

	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

type FoodRatingCreateRepo interface {
	CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error
}

type FoodRatingCreateBiz struct {
	repo FoodRatingCreateRepo
}

func NewFoodRatingCreateBiz(store FoodRatingCreateRepo) *FoodRatingCreateBiz {
	return &FoodRatingCreateBiz{repo: store}
}

func (b *FoodRatingCreateBiz) CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error {
	return b.repo.CreateFoodRating(ctx, data)
}
