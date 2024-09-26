package restaurantbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

type GetListRestaurantStore interface {
	GetList(context context.Context, filter *restaurentmodel.Filter, paging *common.Paging, more ...string) (result []restaurentmodel.Restaurant, err error)
}

type LikeRestaurantStore interface {
	GetRestautantsLike(context context.Context, ids []int) (map[int]int, error)
}

type RestaurantStore struct {
	Store     GetListRestaurantStore
	LikeStore LikeRestaurantStore
}

func NewGetListRestaurantStore(store GetListRestaurantStore, likeStore LikeRestaurantStore) *RestaurantStore {
	return &RestaurantStore{Store: store, LikeStore: likeStore}

}

func (store *RestaurantStore) GetDataByCondition(context context.Context, filter *restaurentmodel.Filter, paging *common.Paging) (result []restaurentmodel.Restaurant, err error) {
	//business logic
	// call db
	result, err = store.Store.GetList(context, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))

	for i, item := range result {
		ids[i] = item.Id
	}

	// get like
	mapList, err := store.LikeStore.GetRestautantsLike(context, ids)
	if err != nil {
		return result, nil

	}
	for i, item := range result {
		result[i].CountLike = mapList[item.Id]
	}

	return result, nil
}
