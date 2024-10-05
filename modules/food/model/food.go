package foodmodel

import (
	"errors"

	"github.com/lehau17/food_delivery/common"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

const (
	EntityName = "foods"
)

type Food struct {
	common.SqlModel
	Name         string                      `json:"name" gorm:"column:name"`
	Description  string                      `json:"description" gorm:"column:description"`
	Price        float64                     `json:"price" gorm:"column:price"`
	Images       *common.Images              `json:"images" gorm:"column:images"`
	CategoryId   int                         `json:"-" gorm:"column:category_id"` // Fixed gorm tag
	Category     *categorymodel.Category     `json:"category" gorm:"preload:false"`
	RestaurantId int                         `json:"-" gorm:"column:restaurant_id"` // Fixed gorm tag
	Restaurant   *restaurentmodel.Restaurant `json:"restaurant" gorm:"preload:false"`
}

func (f *Food) TableName() string { return EntityName }

type FoodCreate struct {
	Name         string         `json:"name" gorm:"column:name"`
	Description  string         `json:"description" gorm:"column:description"`
	Price        float64        `json:"price" gorm:"column:price"`
	Images       *common.Images `json:"images" gorm:"column:images"`
	CategoryId   int            `json:"category_id" gorm:"column:category_id"` // Fixed gorm tag
	RestaurantId int            `json:"-" gorm:"column:restaurant_id"`         // Fixed gorm tag
}

func (f *FoodCreate) TableName() string { return EntityName }

var (
	ErrFoodNotFound       = common.NewFullErrorResponse(400, errors.New("not found food"), "Not found food", "Not found food", "ErrFoodNotFound")
	ErrFoodHasBeenDeleted = common.NewFullErrorResponse(400, errors.New("food has been deleted"), "food has been deleted", "food has been deleted", "ErrFoodHasBeenDeleted")
)

func (f *Food) Mask() {
	f.GenUid(common.DB_FOOD_TYPE)
	f.Category.GenUid(common.DB_CATEGORY_TYPE)
	f.Restaurant.GenUid(common.DB_RESTAURANT_TYPE)
}
