package restaurantratingbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
)

type RestaurantRatingListFoodRepo interface {
	ListFoodRestaurantRating(ctx context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter, paging *common.PagingCursor, moreField ...string) ([]restaurantratingmodel.RestaurantRating, error)
}

type RestaurantRatingListFoodBiz struct {
	repo RestaurantRatingListFoodRepo
}

func NewRestaurantRatingListFoodBiz(repo RestaurantRatingListFoodRepo) *RestaurantRatingListFoodBiz {
	return &RestaurantRatingListFoodBiz{repo: repo}
}

func (s *RestaurantRatingListFoodBiz) ListFoodFoodRating(ctx context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter, paging *common.PagingCursor, moreField ...string) ([]restaurantratingmodel.RestaurantRating, error) {
	return s.repo.ListFoodRestaurantRating(ctx, conditions, filter, paging, moreField...)
}
