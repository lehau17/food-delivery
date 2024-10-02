package userlikerepo

import (
	"context"
	"log"

	"github.com/lehau17/food_delivery/common"
	"github.com/lehau17/food_delivery/components/pubsub"
	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

type UnLikeStore interface {
	FindLike(context context.Context, userId int, resId int) (*userlikerestaurantmodel.Like, error)
	DeleteLike(ctx context.Context, userId int, restaurantId int) error
}

type UnLikeRepo struct {
	store UnLikeStore
	ps    pubsub.PubSub
}

func NewUnlikeRepo(store UnLikeStore, ps pubsub.PubSub) *UnLikeRepo {
	return &UnLikeRepo{ps: ps, store: store}
}

func (ur *UnLikeRepo) UnlikeRestaurant(ctx context.Context, data *userlikerestaurantmodel.Like) error {
	//find like
	foundLike, _ := ur.store.FindLike(ctx, data.UserId, data.RestaurantId)
	if foundLike == nil {
		return userlikerestaurantmodel.ErrLikeNotExist
	}
	//delete like
	if err := ur.store.DeleteLike(ctx, data.UserId, data.RestaurantId); err != nil {
		return err
	}

	//publish 1 message des like in model res
	if err := ur.ps.Publish(ctx, common.TopicUserUnLikeRestaurant, pubsub.NewMessage(common.TopicUserUnLikeRestaurant, data)); err != nil {
		log.Println("Error when des in res like:>>>>", err)
	}

	return nil
}
