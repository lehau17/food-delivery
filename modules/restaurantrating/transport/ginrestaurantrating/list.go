package ginrestaurantrating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	restaurantratingbiz "github.com/lehau17/food_delivery/modules/restaurantrating/biz"
	restaurantratingrepo "github.com/lehau17/food_delivery/modules/restaurantrating/repo"
	restaurantratingstorage "github.com/lehau17/food_delivery/modules/restaurantrating/storage"
)

func ListRestaurantRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantRatingIdStr := c.Param("id")
		restaurantRatingUid, err := common.Decode(restaurantRatingIdStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		// var paging common.PagingCursor
		// if err := c.ShouldBind(&paging); err != nil {
		// 	panic(common.ErrInvalidRequest(err))
		// }
		// paging.FullFill()
		// var filter restaurantratingmodel.Filter
		// if err := c.ShouldBind(&filter); err != nil {
		// 	panic(common.ErrInvalidRequest(err))
		// }
		// filter.FullFill()
		// u := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := restaurantratingrepo.NewRestaurantRatngFindRepo(store)
		biz := restaurantratingbiz.NewRestaurantRatingFindBiz(repo)
		data, err := biz.FindFoodRating(c.Request.Context(), map[string]interface{}{"id": restaurantRatingUid.GetLocalId()}, "User")
		if err != nil {
			panic(err)
		}
		data.Mask()
		c.JSON(http.StatusOK, common.SimplyAppResponse(data))
	}
}
