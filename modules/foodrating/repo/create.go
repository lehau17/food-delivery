package foodratingrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

type FoodRatingCreateStore interface {
	CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error
	FindFoodRating(
		ctx context.Context,
		conditions map[string]interface{},
		morePreload ...string) (*foodratingmodel.FoodRating, error)
}

type FoodRatingCreateRepo struct {
	store FoodRatingCreateStore
}

func NewFoodRatingCreateRepo(store FoodRatingCreateStore) *FoodRatingCreateRepo {
	return &FoodRatingCreateRepo{store: store}
}

func (r *FoodRatingCreateRepo) CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error {
	foundFoodRating, err := r.store.FindFoodRating(ctx, map[string]interface{}{"user_id": data.UserId, "food_id": data.FoodId})
	if err != nil {
		if appErr, ok := err.(*common.AppError); ok {
			if appErr.Key != "ErrRecordNotFound" {
				return err
			}
		} else {
			return err

		}
	}
	if foundFoodRating != nil {
		return foodratingmodel.ErrFoodRatingExists
	}
	return r.store.CreateFoodRating(ctx, data)
}
