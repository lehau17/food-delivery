package likefoodrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	"github.com/lehau17/food_delivery/components/pubsub"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
)

type LikeFoodDeleteStore interface {
	FindLike(ctx context.Context, foodId int, userId int) (*likefoodmodel.LikeFood, error)
	DeleteLike(ctx context.Context, foodId int, userId int) error
}

type LikeFoodDeleteRepo struct {
	store LikeFoodDeleteStore
	ps    pubsub.PubSub
}

func NewLikeFoodDeleteRepo(store LikeFoodDeleteStore, ps pubsub.PubSub) *LikeFoodDeleteRepo {
	return &LikeFoodDeleteRepo{store: store, ps: ps}
}

func (r *LikeFoodDeleteRepo) DeleteLikeFood(ctx context.Context, data *likefoodmodel.LikeFoodCreate) error {
	likeFound, err := r.store.FindLike(ctx, data.FoodId, data.UserId)
	if err != nil {
		if appError, ok := err.(*common.AppError); ok {
			if appError.Key == "ErrRecordNotFound" {
				return err
			}
		} else {

			return err
		}
	}
	if likeFound == nil {
		return likefoodmodel.ErrLikeNotExists
	}
	r.ps.Publish(ctx, common.TopicUserUnLikeFood, pubsub.NewMessage(common.TopicUserUnLikeFood, data))
	return r.store.DeleteLike(ctx, data.FoodId, data.UserId)
}
