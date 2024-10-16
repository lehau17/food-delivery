package userrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	mailProvider "github.com/lehau17/food_delivery/components/mailprovider"
	"github.com/lehau17/food_delivery/components/pubsub"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
)

type UserForgotPasswordStore interface {
	// ChangePassword(ctx context.Context, data *usermodel.ChangePassword) error
	Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	ForgotPassword(ctx context.Context, data *usermodel.UserSetVerifyForgotPassword) (string, error)
}

type genToken interface {
	Generate() string
}

type userForgotPasswordRepo struct {
	store UserForgotPasswordStore
	ps    pubsub.PubSub
	gen   genToken
}

func NewUserForgotPasswordRepo(store UserForgotPasswordStore, ps pubsub.PubSub, gen genToken) *userForgotPasswordRepo {
	return &userForgotPasswordRepo{store: store, ps: ps, gen: gen}
}

type dataMail struct {
	Email string
	Url   string
}

func (r *userForgotPasswordRepo) ForgotPassword(ctx context.Context, data *usermodel.UserForgotPassword) error {
	foundUser, err := r.store.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return err
	}
	if foundUser == nil {
		return usermodel.ErrUserNotExists
	}
	if foundUser.Status == 0 {
		return usermodel.ErrUserDisable
	}
	Url, err := r.store.ForgotPassword(ctx, &usermodel.UserSetVerifyForgotPassword{Email: data.Email, Token: r.gen.Generate()})
	if err != nil {
		return err
	}
	if Url == "" {
		return err
	}

	payload := dataMail{Email: data.Email, Url: Url}
	r.ps.Publish(ctx,
		common.TopicSendMailWhileUserForgotPassword,
		pubsub.NewMessage(common.TopicSendMailWhileUserForgotPassword,
			&mailProvider.DataSendMail{To: []string{data.Email}, Subject: "Quên mật khẩu", TemplateName: "forgot_password.html", Payload: payload}))
	return nil
}
