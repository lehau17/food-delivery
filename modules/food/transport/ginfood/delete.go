package ginfood

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodbiz "github.com/lehau17/food_delivery/modules/food/biz"
	foodrepo "github.com/lehau17/food_delivery/modules/food/repo"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
)

func DeleteFood(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		foodStore := foodstorage.NewSqlStore(db)
		repo := foodrepo.NewFoodDeleteRepo(foodStore, appCtx.GetPubSub())
		biz := foodbiz.NewFoodDeleteRepo(repo)
		if err := biz.DeleteFood(context.Background(), id, u.GetUId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Deleted food"}))

	}
}
