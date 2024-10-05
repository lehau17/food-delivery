package foodrepo

import (
	"context"
	"errors"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodCheckOwner interface {
	CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
}

type FoodCheckOwnerRepo struct {
	FoodStore FoodCheckOwner
}

func NewFoodCheckOwnerRepo(foodStore FoodCheckOwner) *FoodCheckOwnerRepo {
	return &FoodCheckOwnerRepo{FoodStore: foodStore}
}
func (b *FoodCreateRepo) CheckOwner(ctx context.Context, data *foodmodel.FoodCreate, userId int) error {
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
