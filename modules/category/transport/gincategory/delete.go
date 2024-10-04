package gincategory

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	categortbiz "github.com/lehau17/food_delivery/modules/category/biz"
	categorystorage "github.com/lehau17/food_delivery/modules/category/storage"
)

func DeleteCategory(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		store := categorystorage.NewSqlStore(db)
		biz := categortbiz.NewCategoryDeleteBiz(store)

		if err := biz.DeleteCategory(context.Background(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Delete category"}))
	}
}
