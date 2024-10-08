package foodratingbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

type FoodRatngUpdateStore interface {
	UpdateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingUpdate, oldData *foodratingmodel.FoodRating) error
	FindFoodRating(ctx context.Context,
		conditions map[string]interface{},
		morePreload ...string) (*foodratingmodel.FoodRating, error)
}

type FoodRatingUpdateBiz struct {
	store FoodRatngUpdateStore
}

func NewFoodRatingUpdateBiz(store FoodRatngUpdateStore) *FoodRatingUpdateBiz {
	return &FoodRatingUpdateBiz{store: store}
}

func (b *FoodRatingUpdateBiz) UpdateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingUpdate) error {
	foundFoodRating, err := b.store.FindFoodRating(ctx, map[string]interface{}{"id": data.Id, "status": 1})
	if err != nil {
		return err
	}
	if foundFoodRating == nil {
		return foodratingmodel.ErrFoodRatingNotExists
	}
	if foundFoodRating.UserId != data.UserId {
		return common.ErrPermission()
	}
	return b.store.UpdateFoodRating(ctx, data, foundFoodRating)
}
