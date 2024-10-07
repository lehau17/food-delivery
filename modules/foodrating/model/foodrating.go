package foodratingmodel

import (
	"errors"

	"github.com/lehau17/food_delivery/common"
)

const (
	EntityName = "food_ratings"
)

type FoodRating struct {
	common.SqlModel
	Point   float32      `json:"point" gorm:"column:point"`
	Comment string       `json:"comment" gorm:"column:comment"`
	UserId  int          `json:"user_id" gorm:"column:user_id"`
	User    *common.User `json:"user" gorm:"preload:false"`
	FoodId  int          `json:"food_id" gorm:"column:food_id"`
}

func (fr *FoodRating) TableName() string {
	return EntityName
}

type FoodRatingCreate struct {
	Point   float32 `json:"point" gorm:"column:point" binding:"required"`
	Comment string  `json:"comment" gorm:"column:comment" binding:"required"`
	UserId  int     `json:"-" gorm:"column:user_id"`
	FoodId  int     `json:"-" gorm:"column:food_id"`
}

func (fr *FoodRatingCreate) TableName() string {
	return EntityName
}

var (
	ErrFoodRatingExists    = common.NewCustomError(errors.New("food rating exists"), "Food Rating Exists", "ErrFoodRatingExists")
	ErrFoodRatingNotExists = common.NewCustomError(errors.New("food rating not exists"), "Food Rating not Exists", "ErrFoodRatingNotExists")
)
