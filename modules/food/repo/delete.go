package foodrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodDeleteStore interface {
	DeleteFood(ctx context.Context, id int) error
	// CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
	FindFoodWithRestaurant(ctx context.Context, foodId int, userId int) (*foodmodel.Food, error)
}

type FoodDeleteRepo struct {
	FoodStore FoodDeleteStore
}

func NewFoodDeleteRepo(foodStore FoodDeleteStore) *FoodDeleteRepo {
	return &FoodDeleteRepo{FoodStore: foodStore}
}

func (b *FoodDeleteRepo) DeleteFood(ctx context.Context, id int, userId int) error {
	food, err := b.FoodStore.FindFoodWithRestaurant(ctx, id, userId)
	if err != nil {
		return err
	}
	if food == nil || food.Id == 0 {
		return foodmodel.ErrFoodNotFound
	}
	if food.Status == 0 {
		return foodmodel.ErrFoodHasBeenDeleted
	}
	if err := b.FoodStore.DeleteFood(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity("food", err)
	}
	return nil
}
