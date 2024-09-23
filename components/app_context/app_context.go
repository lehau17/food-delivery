package appcontext

import (
	uploadprovider "github.com/lehau17/food_delivery/components/provider"
	"gorm.io/gorm"
)

type AppContect interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() *uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider *uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadProvider *uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() *uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
