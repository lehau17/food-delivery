package restaurantratingrepo

import (
	"context"
)

type RestaurantRatingDeleteStore interface {
	DeleteRestaurantRating(ctx context.Context, id int) error
}

type RestaurantRatngDeleteRepo struct {
	store RestaurantRatingDeleteStore
}

func NewRestaurantRatngDeleteRepo(store RestaurantRatingDeleteStore) *RestaurantRatngDeleteRepo {
	return &RestaurantRatngDeleteRepo{store: store}
}

func (r *RestaurantRatngDeleteRepo) DeleteRestaurantRating(ctx context.Context, id int) error {
	return r.store.DeleteRestaurantRating(ctx, id)
}
