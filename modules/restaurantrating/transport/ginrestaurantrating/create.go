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

func CreateRestaurantRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantIdStr := c.Param("id")
		restaurantUid, err := common.Decode(restaurantIdStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data restaurantratingmodel.RestaurantRatingCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		u := c.MustGet(common.CurrentUser).(common.Requester)
		data.RestaurantId = int(restaurantUid.GetLocalId())
		data.UserId = u.GetUId()
		store := restaurantratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := restaurantratingrepo.NewRestaurantRatngCreateRepo(store)
		biz := restaurantratingbiz.NewRestaurantRatingCreateBiz(repo)
		if err := biz.CreateFoodRating(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Created Food Rating"}))
	}
}
