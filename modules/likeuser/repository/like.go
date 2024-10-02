package userlikerepo

import (
	"context"
	"log"

	"github.com/lehau17/food_delivery/common"
	"github.com/lehau17/food_delivery/components/pubsub"
	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

type LikeStore interface {
	CreateLike(ctx context.Context, data *userlikerestaurantmodel.Like) error
	FindLike(context context.Context, userId int, resId int) (*userlikerestaurantmodel.Like, error)
}

type LikeRepo struct {
	store LikeStore
	ps    pubsub.PubSub
}

func NewLikeRepo(store LikeStore, pubsub pubsub.PubSub) *LikeRepo {
	return &LikeRepo{store: store, ps: pubsub}
}

func (likeRepo *LikeRepo) LikeRestautant(ctx context.Context, data *userlikerestaurantmodel.Like) error {
	// find like exist
	foundLike, _ := likeRepo.store.FindLike(ctx, data.UserId, data.RestaurantId)
	if foundLike != nil {
		return userlikerestaurantmodel.ErrLikeExist
	}

	if err := likeRepo.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(common.TopicUserLikeRestaurant, data)); err != nil {
		log.Println("Error while like res>>>>>", err)
	}

	return likeRepo.store.CreateLike(ctx, data)
}
