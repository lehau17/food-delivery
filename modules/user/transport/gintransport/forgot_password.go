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

func ForgotPassword(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserForgotPassword
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSqlStore(appCtx.GetMainDBConnection(), appCtx.GetRedis())
		repo := userrepo.NewUserForgotPasswordRepo(store, appCtx.GetPubSub())
		biz := userbiz.NewUserForgotPasswordBiz(repo)
		if err := biz.ForgotPassword(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Chúng tôi đã gửi Email xác nhận đến tài khoản. Vui lòng kiểm tra Email"}))
	}
}
