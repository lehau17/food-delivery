package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
)

func CheckRole(act appcontext.AppContect, role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		flag := false
		for i := range role {
			if role[i] == u.GetRole() {
				flag = true
				break
			}
		}
		if !flag {
			panic(common.ErrPermissipn())
		}
		c.Next()
	}
}
