package ginLikefood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	likefoodbiz "github.com/lehau17/food_delivery/modules/likefood/biz"
	likefoodrepo "github.com/lehau17/food_delivery/modules/likefood/repo"
	likefoodstorage "github.com/lehau17/food_delivery/modules/likefood/storage"
	userstorage "github.com/lehau17/food_delivery/modules/user/storage"
)

func GetListUserLikeFood(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId := c.Param("id")
		uid, err := common.Decode(foodId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		likeFoodStore := likefoodstorage.NewSqlStore(db)
		userStore := userstorage.NewSqlStore(db)
		repo := likefoodrepo.NewLikeFoodGetListUserRepo(userStore, likeFoodStore)
		biz := likefoodbiz.NewLikeFoodListUserBiz(repo)
		data, err := biz.GetListUserLikeFood(c.Request.Context(), int(uid.GetLocalId()))
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(data))
	}
}
