package restaurantratingrepo

import (
	"context"

	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

type RestaurantRatingFindStore interface {
	FindRestaurantRating(ctx context.Context, conditions map[string]interface{}, moreField ...string) (*restaurantratingmodel.RestaurantRating, error)
}

type RestaurantRatngFindRepo struct {
	store RestaurantRatingFindStore
}

func NewRestaurantRatngFindRepo(store RestaurantRatingFindStore) *RestaurantRatngFindRepo {
	return &RestaurantRatngFindRepo{store: store}
}

func (r *RestaurantRatngFindRepo) FindRestaurantRating(ctx context.Context, conditions map[string]interface{}, moreField ...string) (*restaurantratingmodel.RestaurantRating, error) {
	return r.store.FindRestaurantRating(ctx, conditions, moreField...)
}
