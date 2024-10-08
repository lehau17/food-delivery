package ginfoodrating

import (
	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodratingbiz "github.com/lehau17/food_delivery/modules/foodrating/biz"
	foodratingmodel "github.com/lehau17/food_delivery/modules/foodrating/model"
	foodratingstorage "github.com/lehau17/food_delivery/modules/foodrating/storage"
)

func GetListFoodRating(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter foodratingmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.FullFill()
		var paging common.PagingCursor
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		foodIdStr := c.Param("id")
		foodUid, err := common.Decode(foodIdStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := foodratingstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := foodratingbiz.NewFoodRatingListBiz(store)
		data, err := biz.GetListFoodRating(c.Request.Context(), map[string]interface{}{"food_id": foodUid.GetLocalId()}, &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range data {
			data[i].Mask()
		}
		dataLen := len(data)
		if dataLen > 0 && dataLen == paging.Limit {
			paging.Cursor = data[len(data)-1].Fake_id.String()
		}
		c.JSON(200, common.NewAppResponse(data, paging, filter))
	}
}
