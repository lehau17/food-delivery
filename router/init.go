package router

import (
	"github.com/gin-gonic/gin"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/middlewares"
)

func InitRouter(mode string, appCtx appcontext.AppContect) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(middlewares.Recovery(appCtx))
	InitUserRouter(r, appCtx)
	InitResRouter(r, appCtx)
	InitCategoryRouter(r, appCtx)
	InitFoodRatingRouter(r, appCtx)
	InitFoodRouter(r, appCtx)
	return r
}
