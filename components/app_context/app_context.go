package appcontext

import (
	uploadprovider "github.com/lehau17/food_delivery/components/provider"
	"github.com/lehau17/food_delivery/components/pubsub"
	"gorm.io/gorm"
)

type AppContect interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() *uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider *uploadprovider.UploadProvider
	secret         string
	ps             pubsub.PubSub
}

func NewAppContext(db *gorm.DB, uploadProvider *uploadprovider.UploadProvider, secret string, ps pubsub.PubSub) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secret: secret, ps: ps}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() *uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secret
}
func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.ps
}
