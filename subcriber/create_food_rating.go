package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	foodrestaurantmodel "github.com/lehau17/food_delivery/modules/food_restaurant/model"
	foodrestaurantstorage "github.com/lehau17/food_delivery/modules/food_restaurant/storage"
)

// type data struct {
// 	FoodId       int `json:"food_id"  gorm:"column:food_id"`
// 	RestaurantId int `json:"restaurant_id"  gorm:"column:restaurant_id"`
// }

func CreateFoodRestaurantWhileCreateFood(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Create food rating",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := foodrestaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			foodRestaurant := message.Data().(foodrestaurantmodel.FoodsRestaurantCreate)

			return store.CreateRestaurantFood(context.Background(), &foodRestaurant)
		},
	}
}
