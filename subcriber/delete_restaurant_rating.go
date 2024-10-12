package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	foodrestaurantstorage "github.com/lehau17/food_delivery/modules/food_restaurant/storage"
)

type dataInterface interface {
	GetFoodIdModel() int
	GetResIdModel() int
}

func DeleteFood_RestaurantWhileDeleteFood(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Delete food rating",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := foodrestaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			dataInterface := message.Data().(dataInterface)
			return store.DeleteRestaurantFood(context.Background(), dataInterface.GetResIdModel(), dataInterface.GetFoodIdModel())
		},
	}
}
