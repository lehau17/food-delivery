package restaurantratingbiz

import (
	"context"

	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

type RestaurantRatingFindRepo interface {
	FindRestaurantRating(ctx context.Context, conditions map[string]interface{}, moreField ...string) (*restaurantratingmodel.RestaurantRating, error)
}

type RestaurantRatingFindBiz struct {
	repo RestaurantRatingFindRepo
}

func NewRestaurantRatingFindBiz(repo RestaurantRatingFindRepo) *RestaurantRatingFindBiz {
	return &RestaurantRatingFindBiz{repo: repo}
}

func (s *RestaurantRatingFindBiz) FindFoodRating(ctx context.Context, conditions map[string]interface{}, moreField ...string) (*restaurantratingmodel.RestaurantRating, error) {
	return s.repo.FindRestaurantRating(ctx, conditions, moreField...)
}
