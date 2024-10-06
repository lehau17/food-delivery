package likefoodbiz

import (
	"context"

	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

type LikeFoodRepo interface {
	CreateLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error
}

type LikeFoodCreateBiz struct {
	repo LikeFoodRepo
}

func NewLikeFoodCreateBiz(repo LikeFoodRepo) *LikeFoodCreateBiz {
	return &LikeFoodCreateBiz{repo: repo}
}

func (b *LikeFoodCreateBiz) Create(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error {
	return b.repo.CreateLikeFood(ctx, data)
}
