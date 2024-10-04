package foodrepo

import (
	"context"
	"errors"

	"github.com/lehau17/food_delivery/common"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

type FoodCreateStore interface {
	CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
}

type CategoryStore interface {
	FindCategory(ctx context.Context, conditions map[string]interface{}) (*categorymodel.Category, error)
}

type RestaurantStore interface {
	FindRestaurant(ctx context.Context, conditions map[string]interface{}) (*restaurentmodel.Restaurant, error)
}

type FoodCreateRepo struct {
	FoodStore FoodCreateStore
	CateStore CategoryStore
	ResStore  RestaurantStore
}

func NewFoodCreateRepo(foodStore FoodCreateStore, cateStore CategoryStore, resStore RestaurantStore) *FoodCreateRepo {
	return &FoodCreateRepo{FoodStore: foodStore, CateStore: cateStore, ResStore: resStore}
}

func (b *FoodCreateRepo) CreateFood(ctx context.Context, data *foodmodel.FoodCreate, userId int) error {
	foundRestaurant, err := b.ResStore.FindRestaurant(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return err
	}
	if foundRestaurant == nil || foundRestaurant.Id == 0 {
		return common.ErrRecordNotFound(errors.New("restaurant not found"))
	}
	foundCate, errCate := b.CateStore.FindCategory(ctx, map[string]interface{}{"id": data.CategoryId})
	if errCate != nil {
		return errCate
	}
	if foundCate == nil || foundCate.Id == 0 {
		return common.ErrRecordNotFound(errors.New("category not found"))
	}
	data.RestaurantId = foundRestaurant.Id

	return b.FoodStore.CreateFood(ctx, data)
}
