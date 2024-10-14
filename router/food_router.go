package router

import (
	"github.com/gin-gonic/gin"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/middlewares"
	"github.com/lehau17/food_delivery/modules/food/transport/ginfood"
	"github.com/lehau17/food_delivery/modules/foodrating/transport/ginfoodrating"
	ginLikefood "github.com/lehau17/food_delivery/modules/likefood/transport/gintranport"
)

func InitFoodRouter(r *gin.Engine, ctx appcontext.AppContect) {
	gFood := r.Group("/foods", middlewares.CheckAuth(ctx))
	gFood.POST("/", ginfood.CreateFood(ctx))
	gFood.POST("/:id/like", ginLikefood.LikeFood(ctx))
	gFood.POST("/:id/rating", ginfoodrating.CreateFoodRating(ctx))
	gFood.GET("/:id/rating", ginfoodrating.GetListFoodRating(ctx))
	gFood.POST("/:id/rating/delete", ginfoodrating.DeleteFoodRating(ctx))
	gFood.POST("/:id/unlike", ginLikefood.UnlikeFood(ctx))
	gFood.GET("/:id", ginfood.FindFood(ctx))
	gFood.GET("/:id/users-like", ginLikefood.GetListUserLikeFood(ctx))
	gFood.GET("/", ginfood.GetList(ctx))
	//
	gFood.PATCH("/:id", ginfood.UpdateFood(ctx))
	//
	gFood.DELETE("/:id", ginfood.DeleteFood(ctx))
}
