package foodmodel

import (
	"errors"
	"log"

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
	Like         int                         `json:"like" gorm:"column:like"`
	AvgPoint     float32                     `json:"avg_point" gorm:"column:avg_point"`
}

func (f *Food) TableName() string { return EntityName }

type FoodCreate struct {
	ID           int            `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name         string         `json:"name" binding:"required" gorm:"column:name"`
	Description  string         `json:"description" gorm:"column:description"`
	Price        float64        `json:"price" binding:"required" gorm:"column:price"`
	Images       *common.Images `json:"images" gorm:"column:images"`
	CategoryId   int            `json:"category_id" binding:"required" gorm:"column:category_id"`
	RestaurantId int            `json:"-" gorm:"column:restaurant_id"` // Fixed gorm tag
	AvgPoint     float32        `json:"-" gorm:"column:avg_point"`
}

type FoodUpdate struct {
	Name         *string        `json:"name,omitempty" gorm:"column:name"`
	Description  *string        `json:"description,omitempty" gorm:"column:description"`
	Price        *float64       `json:"price,omitempty" gorm:"column:price"`
	Images       *common.Images `json:"images,omitempty" gorm:"column:images"`
	CategoryId   *int           `json:"category_id,omitempty" gorm:"column:category_id"`
	RestaurantId *int           `json:"-" gorm:"column:restaurant_id"`
}

func (f *FoodCreate) TableName() string { return EntityName }
func (f *FoodUpdate) TableName() string { return EntityName }

var (
	ErrFoodNotFound       = common.NewFullErrorResponse(400, errors.New("not found food"), "Not found food", "Not found food", "ErrFoodNotFound")
	ErrFoodHasBeenDeleted = common.NewFullErrorResponse(400, errors.New("food has been deleted"), "food has been deleted", "food has been deleted", "ErrFoodHasBeenDeleted")
)

func (f *Food) Mask() {
	f.GenUid(common.DB_FOOD_TYPE)
	log.Println(f.Fake_id)
	if f.Category != nil {
		f.Category.Mask()
	}
	if f.Restaurant != nil {
		f.Restaurant.Mask(false)
	}
}
