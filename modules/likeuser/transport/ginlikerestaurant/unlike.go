package ginlikerestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	likerestaurantbiz "github.com/lehau17/food_delivery/modules/likeuser/biz"
	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
	userlikerepo "github.com/lehau17/food_delivery/modules/likeuser/repository"
	likestorage "github.com/lehau17/food_delivery/modules/likeuser/storage"
)

func UnlikeRestaurant(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		resId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		u := c.MustGet(common.CurrentUser).(common.Requester)
		store := likestorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := userlikerepo.NewUnlikeRepo(store, appCtx.GetPubSub())
		biz := likerestaurantbiz.NewUnlikeBiz(repo)
		if err := biz.UnlikeRestaurant(c.Request.Context(), &userlikerestaurantmodel.Like{RestaurantId: resId, UserId: u.GetUId()}); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(map[string]interface{}{"message": "UnLike restaurant successfully "}))
	}
}
