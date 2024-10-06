package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
)

func DesLikeCountAfterUserDisLikeFood(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Decrease like count after user dislikes food",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := foodstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasFoodId)

			return store.DesLikeFood(ctx, likeData.GetFoodId())
		},
	}
}
