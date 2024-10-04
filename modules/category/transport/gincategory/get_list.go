package gincategory

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	categortbiz "github.com/lehau17/food_delivery/modules/category/biz"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
	categorystorage "github.com/lehau17/food_delivery/modules/category/storage"
)

func GetListCategory(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		store := categorystorage.NewSqlStore(db)
		biz := categortbiz.NewCategoryGetListBiz(store)
		var paging common.PagingCursor
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if paging.Cursor != "" {
			uid, err := common.FromBase58(paging.Cursor)
			log.Println("Uid:>>>", uid.GetLocalId())
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
			paging.RealCursor = int(uid.GetLocalId())
		}
		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()
		data, err := biz.GetList(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		cursor := data[len(data)-1]
		cursor.GenUid(common.DB_CATEGORY_TYPE)
		// cursor.Fake_id = common.NewUid(uint32(cursor.Id), common.DB_CATEGORY_TYPE, 1)
		log.Println("local Id: ", cursor.Fake_id.GetLocalId())
		paging.Cursor = cursor.Fake_id.String()
		c.JSON(http.StatusOK, common.NewAppResponse(data, paging, filter))
	}
}
