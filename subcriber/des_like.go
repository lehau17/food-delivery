package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	restaurentstorage "github.com/lehau17/food_delivery/modules/restaurent/storage"
)

func DesLikeCountAfterUserDisLikeRestaurant(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Decrease like count after user dislikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurentstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.UnlikeRestaurant(ctx, likeData.GetResId())
		},
	}
}
