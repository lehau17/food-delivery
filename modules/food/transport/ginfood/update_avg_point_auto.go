package ginfood

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	foodbiz "github.com/lehau17/food_delivery/modules/food/biz"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
)

func UpdateAvgPointsAuto(appCtx appcontext.AppContect) {

	store := foodstorage.NewSqlStore(appCtx.GetMainDBConnection())
	biz := foodbiz.NewFoodUpdateAvgPointBiz(store)

	if err := biz.FoodUpdateAvgPointBiz(context.Background()); err != nil {
		panic(err)
	}
}
