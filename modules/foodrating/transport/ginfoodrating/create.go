package ginfoodrating

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodratingbiz "github.com/lehau17/food_delivery/modules/foodrating/biz"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
	foodratingrepo "github.com/lehau17/food_delivery/modules/foodrating/repo"
	foodratingstorage "github.com/lehau17/food_delivery/modules/foodrating/storage"
)

// food-rating
func CreateFoodRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		var body foodratingmodel.FoodRatingCreate
		if err := c.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		body.UserId = u.GetUId()
		foodIdStr := c.Param("id")
		if foodIdStr == "" {
			panic(common.ErrInvalidRequest(errors.New("invalid request")))
		}
		foodUid, err := common.Decode(foodIdStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		body.FoodId = int(foodUid.GetLocalId())

		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := foodratingrepo.NewFoodRatingCreateRepo(store)
		biz := foodratingbiz.NewFoodRatingCreateBiz(repo)
		if err := biz.CreateFoodRating(c.Request.Context(), &body); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Created food rating"}))
	}
}
