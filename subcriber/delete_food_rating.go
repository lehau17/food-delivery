package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	foodrestaurantstorage "github.com/lehau17/food_delivery/modules/food_restaurant/storage"
)

type dataDeleted struct {
	FoodId       int `json:"food_id"  gorm:"column:food_id"`
	RestaurantId int `json:"restaurant_id"  gorm:"column:restaurant_id"`
}

func DeleteFoodRestaurantWhileDeleteFood(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Delete food rating",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := foodrestaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			foodId := message.Data().(dataDeleted)
			return store.DeleteRestaurantFood(context.Background(), foodId.RestaurantId, foodId.FoodId)
		},
	}
}
