package likefoodrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

type LikeFoodStore interface {
	FindLike(ctx context.Context, foodId int, userId int) (*likefoodmodel.LikeFood, error)
	CreateLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error
}

type LikeFoodRepo struct {
	store LikeFoodStore
	// ps
}

func NewLikeFoodRepo(store LikeFoodStore) *LikeFoodRepo {
	return &LikeFoodRepo{store: store}
}

func (r *LikeFoodRepo) CreateLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error {
	likeFound, err := r.store.FindLike(ctx, data.FoodId, data.UserId)
	if err != nil {
		if appError, ok := err.(*common.AppError); ok {
			if appError.Key != "ErrRecordNotFound" {
				return err
			}
		} else {

			return err
		}
	}
	if likeFound != nil {
		return likefoodmodel.ErrLikeExists
	}
	return r.store.CreateLikeFood(ctx, data)
}
