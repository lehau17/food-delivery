package likerestaurantbiz

import (
	"context"

	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

type LikeRestaurantRepo interface {
	LikeRestautant(ctx context.Context, data *userlikerestaurantmodel.Like) error
}
type LikeRestaurantBiz struct {
	likerepo LikeRestaurantRepo
}

func NewLikeRestarantBiz(likerepo LikeRestaurantRepo) *LikeRestaurantBiz {
	return &LikeRestaurantBiz{likerepo: likerepo}
}

func (lz *LikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *userlikerestaurantmodel.Like) error {
	if err := lz.likerepo.LikeRestautant(ctx, data); err != nil {
		return err
	}
	return nil
}
