package foodbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodListRepo interface {
	GetListFood(
		ctx context.Context,
		conditions map[string]interface{},
		filter *foodmodel.Filter,
		paging *common.PagingCursor,
		moreLoad ...string) ([]foodmodel.Food, error)
}

type FoodListBiz struct {
	repo FoodListRepo
}

func NewFoodListBiz(repo FoodListRepo) *FoodListBiz {
	return &FoodListBiz{repo: repo}
}

func (b *FoodListBiz) GetList(ctx context.Context,
	conditions map[string]interface{},
	filter *foodmodel.Filter,
	paging *common.PagingCursor) ([]foodmodel.Food, error) {
	return b.repo.GetListFood(ctx, conditions, filter, paging, "Category")
}
