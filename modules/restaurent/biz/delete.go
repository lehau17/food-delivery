package restaurantbiz

import "context"

type DeleteRestaurantStore interface {
	DeleteRestaurant(context context.Context, id int) error
}

type DeleteRestaurantBiz struct {
	Store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *DeleteRestaurantBiz {
	return &DeleteRestaurantBiz{Store: store}
}

func (biz *DeleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	//business logic
	//execution logic
	if err := biz.Store.DeleteRestaurant(context, id); err != nil {
		return err
	}
	return nil
}
