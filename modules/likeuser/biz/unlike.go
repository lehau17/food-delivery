package likerestaurantbiz

import (
	"context"

	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

type UnlikeRepo interface {
	UnlikeRestaurant(ctx context.Context, data *userlikerestaurantmodel.Like) error
}

type UnlikeBiz struct {
	repo UnlikeRepo
}

func NewUnlikeBiz(repo UnlikeRepo) *UnlikeBiz {
	return &UnlikeBiz{repo: repo}
}

func (b *UnlikeBiz) UnlikeRestaurant(ctx context.Context, data *userlikerestaurantmodel.Like) error {
	return b.repo.UnlikeRestaurant(ctx, data)
}
