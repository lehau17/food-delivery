package ginfoodrating

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodratingbiz "github.com/lehau17/food_delivery/modules/foodrating/biz"
	foodratingstorage "github.com/lehau17/food_delivery/modules/foodrating/storage"
)

// food-rating
func DeleteFoodRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		foodRatingStr := c.Param("id")
		if foodRatingStr == "" {
			panic(common.ErrInvalidRequest(errors.New("invalid request")))
		}
		foodRating, err := common.Decode(foodRatingStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		// repo := foodratingrepo.(store)
		biz := foodratingbiz.NewFoodRatingDeleteBiz(store)
		if err := biz.DeleteFoodRating(c.Request.Context(), int(foodRating.GetLocalId()), u.GetUId()); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Deleted food rating"}))
	}
}
