package userlikerepo

import (
	"context"

	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

type LikeStore interface {
	CreateLike(ctx context.Context, data *userlikerestaurantmodel.Like) error
	FindLike(context context.Context, userId int, resId int) (*userlikerestaurantmodel.Like, error)
}

type LikeRepo struct {
	store LikeStore
}

func NewLikeRepo(store LikeStore) *LikeRepo {
	return &LikeRepo{store: store}
}

func (likeRepo *LikeRepo) LikeRestautant(ctx context.Context, data *userlikerestaurantmodel.Like) error {
	// find like exist
	foundLike, err := likeRepo.store.FindLike(ctx, data.UserId, data.RestaurantId)
	if err != nil {
		return err

	}

	if foundLike != nil {
		return userlikerestaurantmodel.ErrLikeExist
	}
	return likeRepo.store.CreateLike(ctx, data)
}
