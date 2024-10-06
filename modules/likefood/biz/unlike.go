package likefoodbiz

import (
	"context"

	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

type LikeFoodDeleteRepo interface {
	DeleteLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error
}

type LikeFoodDeleteBiz struct {
	repo LikeFoodDeleteRepo
}

func NewLikeFoodDeleteBiz(repo LikeFoodDeleteRepo) *LikeFoodDeleteBiz {
	return &LikeFoodDeleteBiz{repo: repo}
}

func (b *LikeFoodDeleteBiz) DeleteLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error {
	return b.repo.DeleteLikeFood(ctx, data)
}
