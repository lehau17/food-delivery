package ginrestaurantrating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	restaurantratingbiz "github.com/lehau17/food_delivery/modules/restaurantrating/biz"
	restaurantratingmodel "github.com/lehau17/food_delivery/modules/restaurantrating/model"
	restaurantratingrepo "github.com/lehau17/food_delivery/modules/restaurantrating/repo"
	restaurantratingstorage "github.com/lehau17/food_delivery/modules/restaurantrating/storage"
)

func FindRestaurantRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantIdStr := c.Param("id")
		restaurantUid, err := common.Decode(restaurantIdStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.PagingCursor
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		var filter restaurantratingmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		// filter.FullFill()
		// u := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := restaurantratingrepo.NewRestaurantRatngListFoodRepo(store)
		biz := restaurantratingbiz.NewRestaurantRatingListFoodBiz(repo)
		data, err := biz.ListRestaurantRating(c.Request.Context(), map[string]interface{}{"restaurant_id": restaurantUid.GetLocalId()}, &filter, &paging, "user")
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask()
		}
		c.JSON(http.StatusOK, common.NewAppResponse(data, paging, filter))
	}
}
