package ginLikefood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
	likefoodbiz "github.com/lehau17/food_delivery/modules/likefood/biz"
	likefoodrepo "github.com/lehau17/food_delivery/modules/likefood/repo"
	likefoodstorage "github.com/lehau17/food_delivery/modules/likefood/storage"
)

func GetListFoodUserLiked(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.PagingCursor
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		var filter foodmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		u := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()
		foodStore := foodstorage.NewSqlStore(db)
		likeFoodStore := likefoodstorage.NewSqlStore(db)
		repo := likefoodrepo.NewLikeFoodGetListFoodRepo(foodStore, likeFoodStore)
		biz := likefoodbiz.NewLikeFoodListFoodUserLikeBiz(repo)
		data, err := biz.GetListFoodUserLiked(c.Request.Context(), u.GetUId(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask()
		}
		c.JSON(http.StatusOK, common.NewAppResponse(data, paging, filter))
	}
}
