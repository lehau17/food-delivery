package restaurentstorage

import (
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
	"gorm.io/gorm"
)

// lưu trữ các phương thức tương tac với database : service
// tạo 1 bộ lưu trữ các phương thức tương tác với database
type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	// Migrate the schema
	db.AutoMigrate(&restaurentmodel.Restaurant{})
	return &sqlStore{db: db}
}
