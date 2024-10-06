package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	foodstorage "github.com/lehau17/food_delivery/modules/food/storage"
)

type HasFoodId interface {
	GetFoodId() int
}

func IncLikeCountAfterUserLikeFood(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Inc like count after user like food",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := foodstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasFoodId)

			return store.IncLikeFood(ctx, likeData.GetFoodId())
		},
	}
}
