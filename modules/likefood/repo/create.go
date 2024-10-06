package likefoodrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	"github.com/lehau17/food_delivery/components/pubsub"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

type LikeFoodStore interface {
	FindLike(ctx context.Context, foodId int, userId int) (*likefoodmodel.LikeFood, error)
	CreateLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error
}

type LikeFoodRepo struct {
	store LikeFoodStore
	ps    pubsub.PubSub
}

func NewLikeFoodRepo(store LikeFoodStore, ps pubsub.PubSub) *LikeFoodRepo {
	return &LikeFoodRepo{store: store, ps: ps}
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
	r.ps.Publish(ctx, common.TopicUserLikeFood, pubsub.NewMessage(common.TopicUserLikeFood, data))
	return r.store.CreateLikeFood(ctx, data)
}
