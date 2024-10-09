package restaurantratingrepo

import (
	"context"

	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

type RestaurantRatingCreateStore interface {
	CreateRestaurantRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error
}

type RestaurantRatngCreateRepo struct {
	store RestaurantRatingCreateStore
}

func NewRestaurantRatngCreateRepo(store RestaurantRatingCreateStore) *RestaurantRatngCreateRepo {
	return &RestaurantRatngCreateRepo{store: store}
}
