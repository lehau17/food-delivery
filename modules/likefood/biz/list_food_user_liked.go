package likefoodbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type LikeFoodGetListFoodUserLikeRepo interface {
	GetListFoodUserLiked(ctx context.Context, userId int, filter *foodmodel.Filter, paging *common.PagingCursor) ([]foodmodel.Food, error)
}

type LikeFoodListFoodUserLikeBiz struct {
	repo LikeFoodGetListFoodUserLikeRepo
}

func NewLikeFoodListFoodUserLikeBiz(repo LikeFoodGetListFoodUserLikeRepo) *LikeFoodListFoodUserLikeBiz {
	return &LikeFoodListFoodUserLikeBiz{repo: repo}
}

func (b *LikeFoodListFoodUserLikeBiz) GetListFoodUserLiked(ctx context.Context, userId int, filter *foodmodel.Filter, paging *common.PagingCursor) ([]foodmodel.Food, error) {
	return b.repo.GetListFoodUserLiked(ctx, userId, filter, paging)
}
