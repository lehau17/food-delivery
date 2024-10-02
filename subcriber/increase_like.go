package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	restaurentstorage "github.com/lehau17/food_delivery/modules/restaurent/storage"
)

type HasRestaurantId interface {
	GetUserId() int
	GetResId() int
}

func IncreaseLikeCountAfterUserDisLikeRestaurant(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Decrease like count after user dislikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurentstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.LikeRestaurant(ctx, likeData.GetResId())
		},
	}
}
