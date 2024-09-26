package usertransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/hasher"
	"github.com/lehau17/food_delivery/components/jwtprovider/jwt"
	userbiz "github.com/lehau17/food_delivery/modules/user/biz"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
	userstorage "github.com/lehau17/food_delivery/modules/user/storage"
)

func Login(atx appcontext.AppContect) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSqlStore(atx.GetMainDBConnection())
		md5 := hasher.NewMd5Hasher()
		jwtProvider := jwt.NewJwtProvider(atx.SecretKey())
		biz := userbiz.NewUserLoginBiz(store, md5, jwtProvider, 30)
		token, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		// foundUser.Mask(false)
		// sign token
		// token, err := biz.
		c.JSON(http.StatusOK, common.SimplyAppResponse(gin.H{"message": "Login success", "data": token}))
	}
}
