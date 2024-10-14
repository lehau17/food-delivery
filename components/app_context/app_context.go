package appcontext

import (
	uploadprovider "github.com/lehau17/food_delivery/components/provider"
	"github.com/lehau17/food_delivery/components/pubsub"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AppContect interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() *uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.PubSub
	GetRedis() *redis.Client
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider *uploadprovider.UploadProvider
	secret         string
	ps             pubsub.PubSub
	rbd            *redis.Client
}

// GetRedis implements AppContect.
func (ctx *appCtx) GetRedis() *redis.Client {
	return ctx.rbd

}

func NewAppContext(db *gorm.DB, uploadProvider *uploadprovider.UploadProvider, secret string, ps pubsub.PubSub, rbd *redis.Client) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secret: secret, ps: ps, rbd: rbd}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

// func (ctx *appCtx) GetRedis() *redis.Client {
// 	return ctx.rbd
// }

func (ctx *appCtx) UploadProvider() *uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secret
}
func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.ps
}
