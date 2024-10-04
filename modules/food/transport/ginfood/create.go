package ginfood

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	categorystorage "github.com/lehau17/food_delivery/modules/category/storage"
	foodbiz "github.com/lehau17/food_delivery/modules/food/biz"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
	foodrepo "github.com/lehau17/food_delivery/modules/food/repo"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
	restaurentstorage "github.com/lehau17/food_delivery/modules/restaurent/storage"
)

func CreateFood(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		var data foodmodel.FoodCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		foodStore := foodstorage.NewSqlStore(db)
		resStore := restaurentstorage.NewSqlStore(db)
		cateStore := categorystorage.NewSqlStore(db)
		repo := foodrepo.NewFoodCreateRepo(foodStore, cateStore, resStore)
		biz := foodbiz.NewFoodCreateBiz(repo)
		if err := biz.CreateFood(context.Background(), &data, u.GetUId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Created food"}))

	}
}
