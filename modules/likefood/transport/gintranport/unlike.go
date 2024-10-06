package ginLikefood

import (
	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	likefoodbiz "github.com/lehau17/food_delivery/modules/likefood/biz"
	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
	likefoodrepo "github.com/lehau17/food_delivery/modules/likefood/repo"
	likefoodstorage "github.com/lehau17/food_delivery/modules/likefood/storage"
)

func UnlikeFood(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		foodId := c.Param("id")
		uid, err := common.Decode(foodId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := likefoodstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := likefoodrepo.NewLikeFoodDeleteRepo(store, appCtx.GetPubSub())
		biz := likefoodbiz.NewLikeFoodDeleteBiz(repo)
		data := likefoodmodel.LikeFoodCreate{FoodId: int(uid.GetLocalId()), UserId: u.GetUId()}
		if err := biz.DeleteLikeFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimplyAppResponse(gin.H{"message": "UnLike food success"}))
	}
}
