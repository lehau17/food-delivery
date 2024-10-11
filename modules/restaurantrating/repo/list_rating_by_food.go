package restaurantratingrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

type RestaurantRatingListFoodStore interface {
	GetListRestaurantRating(ctx context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter, paging *common.PagingCursor, moreField ...string) ([]restaurantratingmodel.RestaurantRating, error)
}

type RestaurantRatngListFoodRepo struct {
	store RestaurantRatingListFoodStore
}

func NewRestaurantRatngListFoodRepo(store RestaurantRatingListFoodStore) *RestaurantRatngListFoodRepo {
	return &RestaurantRatngListFoodRepo{store: store}
}

func (r *RestaurantRatngListFoodRepo) ListFoodRestaurantRating(ctx context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter, paging *common.PagingCursor, moreField ...string) ([]restaurantratingmodel.RestaurantRating, error) {
	return r.store.GetListRestaurantRating(ctx, conditions, filter, paging, moreField...)
}
