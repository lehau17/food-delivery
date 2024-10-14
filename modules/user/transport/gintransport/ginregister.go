package usertransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/hasher"
	userbiz "github.com/lehau17/food_delivery/modules/user/biz"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
	userstorage "github.com/lehau17/food_delivery/modules/user/storage"
)

func RegisterUser(act appcontext.AppContect) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := act.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hasher()
		biz := userbiz.NewUserRegisterBiz(store, md5, act.GetPubSub(), *act.GetRedis())
		if err := biz.RegisterUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Register user successfully registered"}))
	}
}
