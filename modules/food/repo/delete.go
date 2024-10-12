package foodrepo

import (
	"context"
	"log"

	"github.com/lehau17/food_delivery/common"
	"github.com/lehau17/food_delivery/components/pubsub"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodDeleteStore interface {
	DeleteFood(ctx context.Context, id int) error
	// CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
	FindFoodWithRestaurant(ctx context.Context, foodId int, userId int) (*foodmodel.Food, error)
}

type FoodDeleteRepo struct {
	FoodStore FoodDeleteStore
	ps        pubsub.PubSub
}

func NewFoodDeleteRepo(foodStore FoodDeleteStore, ps pubsub.PubSub) *FoodDeleteRepo {
	return &FoodDeleteRepo{FoodStore: foodStore, ps: ps}
}

type DataDeleted struct {
	FoodId       int `json:"food_id"  gorm:"column:food_id"`
	RestaurantId int `json:"restaurant_id"  gorm:"column:restaurant_id"`
}

func (d *DataDeleted) GetFoodId() int {
	return d.FoodId
}

func (d *DataDeleted) GetRestaurantId() int {
	return d.RestaurantId
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
	log.Println("check data food>>>>", food.RestaurantId, food.Id)
	b.ps.Publish(ctx, common.TopicDeleteFoodRestaurant, pubsub.NewMessage(common.TopicDeleteFoodRestaurant, food))
	return nil
}
