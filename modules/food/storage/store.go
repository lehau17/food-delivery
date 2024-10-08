package foodstorage

import (
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	db.AutoMigrate(&foodmodel.Food{})
	return &sqlStore{db: db}
}
