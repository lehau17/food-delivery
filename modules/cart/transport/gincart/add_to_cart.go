package gincart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	cartbiz "github.com/lehau17/food_delivery/modules/cart/biz"
	cartmodel "github.com/lehau17/food_delivery/modules/cart/model"
	cartrepo "github.com/lehau17/food_delivery/modules/cart/repo"
	cartstore "github.com/lehau17/food_delivery/modules/cart/store"
)

func AddToCart(appCtx appcontext.AppContect) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodIdStr := c.Param("id")
		foodUid, err := common.Decode(foodIdStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		u := c.MustGet(common.CurrentUser).(common.Requester)
		var data cartmodel.CartCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.FoodId = int(foodUid.GetLocalId())
		data.UserId = u.GetUId()
		store := cartstore.NewSqlStore(appCtx.GetMainDBConnection(), appCtx.GetRedis())
		repo := cartrepo.NewCartAddToRepo(store)
		biz := cartbiz.NewCartAddToBiz(repo)
		if err := biz.AddToCart(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimplyAppResponse(true))
	}
}
