package foodratingbiz

import "context"

type FoodRatingDeleteStore interface {
	DeleteFoodRating(ctx context.Context, id int, userId int) error
}

type FoodRatingDeleteBiz struct {
	store FoodRatingDeleteStore
}

func NewFoodRatingDeleteBiz(store FoodRatingDeleteStore) *FoodRatingDeleteBiz {
	return &FoodRatingDeleteBiz{
		store: store,
	}
}

func (b *FoodRatingDeleteBiz) DeleteFoodRating(ctx context.Context, id int, userId int) error {
	return b.store.DeleteFoodRating(ctx, id, userId)
}
