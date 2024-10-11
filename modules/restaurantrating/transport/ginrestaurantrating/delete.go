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

func DeleteRestaurantRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodRatingStr := c.Param("id")
		foodRatingUid, err := common.Decode(foodRatingStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// u := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := restaurantratingrepo.NewRestaurantRatngDeleteRepo(store)
		biz := restaurantratingbiz.NewRestaurantRatingDeleteBiz(repo)
		if err := biz.DeleteFoodRating(c.Request.Context(), int(foodRatingUid.GetLocalId())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Deleted Food Rating"}))
	}
}
