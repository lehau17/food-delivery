package restaurantratingbiz

import (
	"context"
)

type RestaurantRatingDeleteRepo interface {
	DeleteRestaurantRating(ctx context.Context, id int) error
}

type RestaurantRatingDeleteBiz struct {
	repo RestaurantRatingDeleteRepo
}

func NewRestaurantRatingDeleteBiz(repo RestaurantRatingDeleteRepo) *RestaurantRatingDeleteBiz {
	return &RestaurantRatingDeleteBiz{repo: repo}
}

func (s *RestaurantRatingDeleteBiz) DeleteFoodRating(ctx context.Context, id int) error {
	return s.repo.DeleteRestaurantRating(ctx, id)
}
