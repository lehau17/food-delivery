package ginrestaurant

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	restaurantbiz "github.com/lehau17/food_delivery/modules/restaurent/biz"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
	restaurentstorage "github.com/lehau17/food_delivery/modules/restaurent/storage"
)

func CreateRestaurant(ctx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()

		var restaurant restaurentmodel.RestaurantCreate
		if err := c.ShouldBind(&restaurant); err != nil {
			// Handle the error and return a response
			// log.
			panic(err)
		}
		u := c.MustGet(common.CurrentUser).(common.Requester)
		restaurant.UserId = u.GetUId()
		// Insert the new restaurant into the database
		store := restaurentstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.Store.CreateRestaurant(c.Request.Context(), &restaurant); err != nil {
			// Handle database insertion error
			// c.JSON(http.StatusInternalServerError, err)
			panic(err)
		}
		restaurant.Mask(false)
		// Success response
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Restaurant created successfully", "data": restaurant}))
	}
}
