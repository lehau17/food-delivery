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

func LikeRestaurant(act appcontext.AppContect) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		u := c.MustGet(common.CurrentUser).(common.Requester)

		// idInt, err := common.FromBase58(id)
		store := likestorage.NewSqlStore(act.GetMainDBConnection())
		repo := userlikerepo.NewLikeRepo(store)
		biz := likerestaurantbiz.NewLikeRestarantBiz(repo)
		if err := biz.LikeRestaurant(c.Request.Context(), &userlikerestaurantmodel.Like{RestaurantId: id, UserId: u.GetUId()}); err != nil {
			panic(err)
		}
		// if err != nil {
		// 	panic(err)
		// }

		c.JSON(http.StatusOK, common.SimplyAppResponse(map[string]interface{}{"message": "Like restaurant successfully "}))
	}
}
