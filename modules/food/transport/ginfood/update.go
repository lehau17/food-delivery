package ginfood

import (
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

func UpdateFood(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data foodmodel.FoodUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		store := foodstorage.NewSqlStore(db)
		repo := foodrepo.NewFoodUpdateRepo(store)
		biz := foodbiz.NewFoodUpdateBiz(repo)
		if err := biz.UpdateFood(c.Request.Context(), &data, id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Updated food"}))
	}
}
