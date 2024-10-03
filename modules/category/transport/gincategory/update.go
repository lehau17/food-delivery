package gincategory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	categortbiz "github.com/lehau17/food_delivery/modules/category/biz"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
	categorystorage "github.com/lehau17/food_delivery/modules/category/storage"
)

func UpdateCategory(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		cateId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := categorystorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := categortbiz.NewCategoryBizUpdate(store)
		if err := biz.UpdateCategory(c.Request.Context(), &data, cateId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Update category Success"}))
	}
}
