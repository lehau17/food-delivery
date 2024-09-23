package restaurantbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

type GetListRestaurantStore interface {
	GetList(context context.Context, filter *restaurentmodel.Filter, paging *common.Paging) (result []restaurentmodel.Restaurant, err error)
}

type RestaurantStore struct {
	Store GetListRestaurantStore
}

func NewGetListRestaurantStore(store GetListRestaurantStore) *RestaurantStore {
	return &RestaurantStore{Store: store}

}

func (store *RestaurantStore) GetDataByCondition(context context.Context, filter *restaurentmodel.Filter, paging *common.Paging) (result []restaurentmodel.Restaurant, err error) {
	//business logic
	// call db
	result, err = store.Store.GetList(context, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
