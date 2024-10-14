package router

import (
	"github.com/gin-gonic/gin"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/middlewares"
	"github.com/lehau17/food_delivery/modules/category/transport/gincategory"
)

func InitCategoryRouter(r *gin.Engine, ctx appcontext.AppContect) {
	gCate := r.Group("/categories", middlewares.CheckAuth(ctx))
	gCate.POST("/", gincategory.CreateCategory(ctx))
	gCate.GET("/", gincategory.GetListCategory(ctx))
	gCate.PATCH("/:id", gincategory.UpdateCategory(ctx))
	gCate.DELETE("/:id", gincategory.DeleteCategory(ctx))
}
