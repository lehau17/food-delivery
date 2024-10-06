package foodrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodListStore interface {
	GetFoods(
		ctx context.Context,
		conditions map[string]interface{},
		filter *foodmodel.Filter,
		paging *common.PagingCursor,
		moreLoad ...string,
	) ([]foodmodel.Food, error)
}

type FoodListRepo struct {
	Store FoodListStore
}

func NewFoodListRepo(store FoodListStore) *FoodListRepo {
	return &FoodListRepo{Store: store}
}

func (r *FoodListRepo) GetListFood(ctx context.Context,
	conditions map[string]interface{},
	filter *foodmodel.Filter,
	paging *common.PagingCursor,
	moreLoad ...string) ([]foodmodel.Food, error) {
	return r.Store.GetFoods(ctx, conditions, filter, paging, moreLoad...)
}
