package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	restaurantbiz "github.com/lehau17/food_delivery/modules/restaurent/biz"
	restaurentstorage "github.com/lehau17/food_delivery/modules/restaurent/storage"
)

func DeleteRestaurant(ctx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db := ctx.GetMainDBConnection()
		store := restaurentstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"isSuccess": 1}))
	}
}
