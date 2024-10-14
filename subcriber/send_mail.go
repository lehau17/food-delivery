package subcriber

import (
	"context"

	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/pubsub"
	mailstorage "github.com/lehau17/food_delivery/modules/mail/storage"
)

type dataSendMail interface {
	GetTo() []string
	GetSubject() string
	GetTemplateName() string
	GetPayLoadSendEmail() interface{}
}

func SendMail(appCtx appcontext.AppContect) cosumerjob {
	return cosumerjob{
		Title: "Send mail",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := mailstorage.NewMailStore()
			data := message.Data().(dataSendMail)
			return store.SendMail(ctx, data.GetTo(), data.GetTemplateName(), data.GetSubject(), data.GetPayLoadSendEmail())
		},
	}
}
