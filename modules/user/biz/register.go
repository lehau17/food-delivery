package userbiz

import (
	"context"
	"time"

	"github.com/lehau17/food_delivery/common"
	"github.com/lehau17/food_delivery/components/pubsub"
	usermodel "github.com/lehau17/food_delivery/modules/user/model"
	"github.com/redis/go-redis/v9"
)

type UserRegisterStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	Find(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type Hash interface {
	Hash(data string) string
}

type userRegisterBiz struct {
	UserRegisterStore UserRegisterStore
	Hash              Hash
	ps                pubsub.PubSub
	rdb               redis.Client
}

func NewUserRegisterBiz(store UserRegisterStore, hash Hash, ps pubsub.PubSub, rdb redis.Client) *userRegisterBiz {
	return &userRegisterBiz{UserRegisterStore: store, Hash: hash, ps: ps, rdb: rdb}
}

type DataSendMail struct {
	to           []string
	subject      string
	templateName string
	payload      interface{}
}

func (d *DataSendMail) GetPayLoadSendEmail() interface{} {
	return d.payload
}

func (d *DataSendMail) GetTo() []string {
	return d.to
}

func (d *DataSendMail) GetSubject() string {
	return d.subject
}

func (d *DataSendMail) GetTemplateName() string {
	return d.templateName
}

type emailData struct {
	Email string
	Otp   string
}

func (b *userRegisterBiz) RegisterUser(ctx context.Context, data *usermodel.UserCreate) error {
	//check user exists
	foundUser, err := b.UserRegisterStore.Find(ctx, map[string]interface{}{"email": data.Email})
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
	data.Role = "user"
	data.Status = 1
	if err := b.UserRegisterStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("users", err)
	}
	payload := emailData{Email: data.Email, Otp: common.GetOtp()}
	b.ps.Publish(ctx, common.TopicSendMailWhileUserCreated,
		pubsub.NewMessage(common.TopicSendMailWhileUserCreated, &DataSendMail{to: []string{data.Email}, subject: "Xác Nhận Tài Khoản", templateName: "email.html", payload: payload}))
	b.rdb.Set(ctx, "otp:"+data.Email, payload.Otp, 300*time.Second)
	return nil

}
