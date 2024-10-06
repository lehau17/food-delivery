package ginfood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodbiz "github.com/lehau17/food_delivery/modules/food/biz"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
	foodrepo "github.com/lehau17/food_delivery/modules/food/repo"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
)

func GetList(appCtx appcontext.AppContect) gin.HandlerFunc {
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
		store := foodstorage.NewSqlStore(appCtx.GetMainDBConnection())
		repo := foodrepo.NewFoodListRepo(store)
		biz := foodbiz.NewFoodListBiz(repo)
		data, err := biz.GetList(c.Request.Context(), nil, &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask()
			// log.Println("Item >>>", item.Fake_id)
		}
		c.JSON(http.StatusOK, common.NewAppResponse(data, paging, filter))
	}
}
