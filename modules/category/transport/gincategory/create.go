package gincategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	categortbiz "github.com/lehau17/food_delivery/modules/category/biz"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
	categorystorage "github.com/lehau17/food_delivery/modules/category/storage"
)

func CreateCategory(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := categorystorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := categortbiz.NewCategoryBizCreate(store)
		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		// go func() {
		// 	for {
		// 		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
		// 			panic(err)
		// 		}
		// 	}
		// }()
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Create category Success"}))
	}
}
