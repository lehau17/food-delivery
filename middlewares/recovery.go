package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
)

func Recovery(ctx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if arrErr, ok := err.(*common.AppError); ok {
					// handle error and response
					c.AbortWithStatusJSON(arrErr.StatusCode, arrErr)
					// xóa stack and return
					panic(err)
				} else {
					appErr := common.ErrInternal(err.(error))
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(appErr)
				}
			}
		}()
		c.Next()
	}
}
