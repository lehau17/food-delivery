package ginrestaurant

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	restaurantbiz "github.com/lehau17/food_delivery/modules/restaurent/biz"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
	restaurentstorage "github.com/lehau17/food_delivery/modules/restaurent/storage"
)

func GetListRestaurant(ctx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var filter restaurentmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			log.Printf("Loi filter")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		filter.Status = []int{1}

		paging.FullFill()
		store := restaurentstorage.NewSqlStore(db)
		biz := restaurantbiz.NewGetListRestaurantStore(store)
		log.Println(paging, filter)
		result, err := biz.GetDataByCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{

				"error": err.Error(),
			})
			return
		}
		//success

		for i := range result {
			result[i].Mask(false)
			result[i].User.Mask(false)
		}
		c.JSON(http.StatusOK, common.NewAppResponse(result, paging, filter))
	}
}
