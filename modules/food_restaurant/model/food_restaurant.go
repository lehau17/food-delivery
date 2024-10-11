package foodrestaurantmodel

import (
	"errors"
	"time"

	"github.com/lehau17/food_delivery/common"
)

const (
	EntityName = "restaurant_foods"
)

type FoodsRestaurant struct {
	FoodId       int        `json:"food_id" gorm:"column:food_id"`
	RestaurantId int        `json:"restaurant_id" gorm:"column:restaurant_id"`
	Status       int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (fr *FoodsRestaurant) TableName() string {
	return EntityName
}

type FoodsRestaurantCreate struct {
	FoodId       int `json:"food_id" binding:"required" gorm:"column:food_id"`
	RestaurantId int `json:"restaurant_id" binding:"required" gorm:"column:restaurant_id"`
}

func (fr *FoodsRestaurantCreate) TableName() string {
	return EntityName
}

var (
	ErrFoodRatingNotExist = common.NewErrorResponse(errors.New("food_rating_not_exist"), "food_rating_not_exist", "food_rating_not_exist", "ErrFoodRatingNotExist")
)
