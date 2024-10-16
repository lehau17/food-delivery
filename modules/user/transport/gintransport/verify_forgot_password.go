package usertransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	userbiz "github.com/lehau17/food_delivery/modules/user/biz"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
	userrepo "github.com/lehau17/food_delivery/modules/user/repo"
	userstorage "github.com/lehau17/food_delivery/modules/user/storage"
)

func SetVerifyForgotPassword(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserSetVerifyForgotPassword
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSqlStore(appCtx.GetMainDBConnection(), appCtx.GetRedis())
		repo := userrepo.NewUserVerifyForgotPasswordRepo(store)
		biz := userbiz.NewUserVerifyForgotPasswordBiz(repo)
		if err := biz.SetVerifyForgotPassword(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(true))
	}
}
