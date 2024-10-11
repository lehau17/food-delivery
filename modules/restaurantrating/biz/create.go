package restaurantratingbiz

import (
	"context"

	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

type RestaurantRatingCreateRepo interface {
	CreateRestaurantRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error
}

type RestaurantRatingCreateBiz struct {
	repo RestaurantRatingCreateRepo
}

func NewRestaurantRatingCreateBiz(repo RestaurantRatingCreateRepo) *RestaurantRatingCreateBiz {
	return &RestaurantRatingCreateBiz{repo: repo}
}

func (s *RestaurantRatingCreateBiz) CreateFoodRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error {
	return s.repo.CreateRestaurantRating(ctx, data)
}
