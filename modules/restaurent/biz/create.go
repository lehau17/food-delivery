package restaurantbiz

import (
	"context"

	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

// create interface mapping storage methods
type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurentmodel.RestaurantCreate) error
}

// struct contains information of interface
type restaurantStore struct {
	Store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *restaurantStore {
	return &restaurantStore{Store: store}
}

func (store *restaurantStore) CreateRestaurant(context context.Context, data *restaurentmodel.RestaurantCreate) error {
	//logic business

	//call to storage
	if err := store.Store.CreateRestaurant(context, data); err != nil {
		return err
	}
	return nil
}
