package router

import (
	"github.com/gin-gonic/gin"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/middlewares"
	ginLikefood "github.com/lehau17/food_delivery/modules/likefood/transport/gintranport"
	usertransport "github.com/lehau17/food_delivery/modules/user/transport/gintransport"
)

func InitUserRouter(r *gin.Engine, ctx appcontext.AppContect) {
	gUser := r.Group("/user")
	gUser.POST("/register", usertransport.RegisterUser(ctx))
	gUser.POST("/register/shipper", usertransport.RegisterUser(ctx))
	gUser.POST("/verify/otp", usertransport.VerifyUserOtp(ctx))
	gUser.POST("/login", usertransport.Login(ctx))
	gUser.POST("/change-password", usertransport.ChangePassword(ctx))
	gUser.POST("/forgot-password", usertransport.ForgotPassword(ctx))
	gUser.GET("/profile", middlewares.CheckAuth(ctx), usertransport.Profile(ctx))
	gUser.GET("/foods-like", middlewares.CheckAuth(ctx), ginLikefood.GetListFoodUserLiked(ctx))
}
