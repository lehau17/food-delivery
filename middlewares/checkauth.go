package middlewares

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/jwtprovider/jwt"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
	userstorage "github.com/lehau17/food_delivery/modules/user/storage"
)

var (
	errWrongAuthHeader = common.NewCustomError(errors.New("wrong auth header"), "wrong auth header", "ErrInvalidRequest")
)

func extractToken(token string) (string, error) {
	parts := strings.Split(token, " ")
	if len(parts) < 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
		return "nil", errWrongAuthHeader
	}
	return parts[1], nil
}

func CheckAuth(act appcontext.AppContect) func(c *gin.Context) {
	jwtProvider := jwt.NewJwtProvider(act.SecretKey())
	return func(c *gin.Context) {
		token, err := extractToken(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		db := act.GetMainDBConnection()
		store := userstorage.NewSqlStore(db, act.GetRedis())
		//decode
		payload, err := jwtProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		id := payload.Uid
		user, err := store.Find(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}
		if user.Status == 0 {
			panic(usermodel.ErrUserDisable)
		}
		user.Mask(false)
		c.Set(common.CurrentUser, user)

		// c.JSON(200, payload)
		c.Next()
	}
}
