package userbiz

import (
	"context"
	"time"

	"github.com/lehau17/food_delivery/common"
	mailProvider "github.com/lehau17/food_delivery/components/mailprovider"
	"github.com/lehau17/food_delivery/components/pubsub"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
	"github.com/redis/go-redis/v9"
)

type RegisterShipperRegisterStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type shipperRegisterBiz struct {
	store RegisterShipperRegisterStore
	Hash  Hash
	ps    pubsub.PubSub
	rdb   redis.Client
}

func NewShipperRegisterBiz(store RegisterShipperRegisterStore, hash Hash, ps pubsub.PubSub, rdb redis.Client) *shipperRegisterBiz {
	return &shipperRegisterBiz{store: store, Hash: hash, ps: ps, rdb: rdb}
}

func (b *shipperRegisterBiz) RegisterShipper(ctx context.Context, data *usermodel.UserCreate) error {
	//check user exists
	foundUser, err := b.store.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		if error, ok := err.(*common.AppError); ok {
			if error.Key != "ErrRecordNotFound" {
				return err
			}
		}
	}
	if foundUser != nil {
		if foundUser.Status == 0 {
			return usermodel.ErrUserDisable
		}
		return usermodel.ErrUserExists
	}

	//create user--
	salt := common.GetSalt(50)
	data.Salt = salt
	hashPassword := b.Hash.Hash(data.Password + salt)
	data.Password = hashPassword
	data.Role = "shipper"
	data.Status = 2
	if err := b.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("users", err)
	}
	payload := emailData{Email: data.Email, Otp: common.GetOtp()}
	b.ps.Publish(ctx, common.TopicSendMailWhileUserCreated,
		pubsub.NewMessage(common.TopicSendMailWhileUserCreated, &mailProvider.DataSendMail{To: []string{data.Email}, Subject: "Xác Nhận Tài Khoản", TemplateName: "email.html", Payload: payload}))
	b.rdb.Set(ctx, "otp:"+data.Email, payload.Otp, 300*time.Second)
	return nil

}
