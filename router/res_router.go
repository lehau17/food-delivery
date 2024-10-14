package router

import (
	"github.com/gin-gonic/gin"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/middlewares"
	"github.com/lehau17/food_delivery/modules/likeuser/transport/ginlikerestaurant"
	"github.com/lehau17/food_delivery/modules/restaurantrating/transport/ginrestaurantrating"
	"github.com/lehau17/food_delivery/modules/restaurent/tranport/ginrestaurant"
	"github.com/lehau17/food_delivery/modules/upload/tranport/ginupload"
)

func InitResRouter(r *gin.Engine, ctx appcontext.AppContect) {
	gRes := r.Group("/restaurants", middlewares.CheckAuth(ctx))
	gRes.GET("/", ginrestaurant.GetListRestaurant(ctx))
	gRes.POST("/", ginrestaurant.CreateRestaurant(ctx))
	gRes.POST("/:id/rating", ginrestaurantrating.CreateRestaurantRating(ctx))
	gRes.GET("/:id/rating", ginrestaurantrating.ListRestaurantRating(ctx))
	gRes.DELETE("/:id/rating", ginrestaurantrating.DeleteRestaurantRating(ctx))
	gRes.DELETE("/:id", ginrestaurant.DeleteRestaurant(ctx))
	gRes.POST("/upload", ginupload.UploadImage(ctx))
	gRes.POST("/:id/like", ginlikerestaurant.LikeRestaurant(ctx))
	gRes.POST("/:id/unlike", ginlikerestaurant.UnlikeRestaurant(ctx))
}
