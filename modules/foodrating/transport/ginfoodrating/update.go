package ginfoodrating

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodratingbiz "github.com/lehau17/food_delivery/modules/foodrating/biz"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
	foodratingstorage "github.com/lehau17/food_delivery/modules/foodrating/storage"
)

// food-rating
func UpdateFoodRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		var body foodratingmodel.FoodRatingUpdate
		if err := c.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		body.UserId = u.GetUId()
		foodRatingIdStr := c.Param("id")
		if foodRatingIdStr == "" {
			panic(common.ErrInvalidRequest(errors.New("invalid request")))
		}
		foodRatingUid, err := common.Decode(foodRatingIdStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		body.Id = int(foodRatingUid.GetLocalId())

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		// repo := foodratingrepo.New(store)
		biz := foodratingbiz.NewFoodRatingUpdateBiz(store)
		if err := biz.UpdateFoodRating(c.Request.Context(), &body); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Updated food rating"}))
	}
}
