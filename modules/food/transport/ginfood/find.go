package ginfood

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodbiz "github.com/lehau17/food_delivery/modules/food/biz"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
	foodrepo "github.com/lehau17/food_delivery/modules/food/repo"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
)

func FindFood(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var filter foodmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.FullFilter()
		log.Println("Check filter:", filter)
		db := appCtx.GetMainDBConnection()
		store := foodstorage.NewSqlStore(db)
		repo := foodrepo.NewFoodFindRepo(store)
		biz := foodbiz.NewFoodFindBiz(repo)

		data, errRes := biz.FindFood(c.Request.Context(), id, &filter)
		if errRes != nil {
			panic(errRes)
		}
		data.Mask()
		log.Println("ID>>>>", data.Fake_id)
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{
			"message": "Fetch food",
			"data":    data,
		}))
	}
}
