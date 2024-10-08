package foodratingbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
)

type FoodRatingListStore interface {
	GetListFoodRating(
		ctx context.Context,
		conditions map[string]interface{},
		filter *foodratingmodel.Filter,
		paging *common.PagingCursor,
		morePreload ...string) ([]foodratingmodel.FoodRating, error)
}

type FoodRatingListBiz struct {
	store FoodRatingListStore
}

func NewFoodRatingListBiz(store FoodRatingListStore) *FoodRatingListBiz {
	return &FoodRatingListBiz{store: store}
}

func (b *FoodRatingListBiz) GetListFoodRating(ctx context.Context,
	conditions map[string]interface{},
	filter *foodratingmodel.Filter,
	paging *common.PagingCursor) ([]foodratingmodel.FoodRating, error) {
	return b.store.GetListFoodRating(ctx, conditions, filter, paging, "User")
}
