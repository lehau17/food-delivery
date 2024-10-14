package router

import (
	"github.com/gin-gonic/gin"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/middlewares"
	"github.com/lehau17/food_delivery/modules/foodrating/transport/ginfoodrating"
)

func InitFoodRatingRouter(r *gin.Engine, ctx appcontext.AppContect) {

	gFoodRating := r.Group("/food-rating", middlewares.CheckAuth(ctx))
	// gFoodRating.POST("/", ginfoodrating.CreateFoodRating(ctx))
	gFoodRating.DELETE("/:id", ginfoodrating.DeleteFoodRating(ctx))
	gFoodRating.PATCH("/:id", ginfoodrating.UpdateFoodRating(ctx))
}
