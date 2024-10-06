package likefoodrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodStore interface {
	GetFoods(ctx context.Context,
		conditions map[string]interface{},
		filter *foodmodel.Filter,
		paging *common.PagingCursor,
		moreLoad ...string) ([]foodmodel.Food, error)
}

type LikeFoodGetListFoodStore interface {
	GetListFoodUserLiked(ctx context.Context, userId int) ([]int, error)
}

type LikeFoodGetListFoodRepo struct {
	foodStore     FoodStore
	likeFoodStore LikeFoodGetListFoodStore
}

func NewLikeFoodGetListFoodRepo(foodStore FoodStore, likeFoodStore LikeFoodGetListFoodStore) *LikeFoodGetListFoodRepo {
	return &LikeFoodGetListFoodRepo{foodStore: foodStore, likeFoodStore: likeFoodStore}
}

func (r *LikeFoodGetListFoodRepo) GetListFoodUserLiked(ctx context.Context, userId int, filter *foodmodel.Filter, paging *common.PagingCursor) ([]foodmodel.Food, error) {
	ids, err := r.likeFoodStore.GetListFoodUserLiked(ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(ids) <= 0 {
		return nil, nil
	}
	foods, err := r.foodStore.GetFoods(ctx, map[string]interface{}{"id": ids}, filter, paging)
	if err != nil {
		return nil, err
	}
	return foods, nil
}
