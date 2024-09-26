package usertransport

import (
	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

func Profile(act appcontext.AppContect) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(*usermodel.User)
		c.JSON(200, common.SimplyAppResponse(user))
	}
}
