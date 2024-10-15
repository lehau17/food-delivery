package cartmodel

import (
	"errors"
	"time"

	"github.com/lehau17/food_delivery/common"
)

const (
	EntityName = "carts"
)

type Cart struct {
	UserId    int          `json:"-" gorm:"user_id"`
	FoodId    int          `json:"food_id" gorm:"food_id"`
	Quantity  int          `json:"quantity" gorm:"quantity"`
	Status    int          `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time   `json:"updated_at" gorm:"column:updated_at"`
	User      *common.User `json:"user" gorm:"preload:false"`
}

func (c *Cart) TableName() string {
	return EntityName
}

type CartCreate struct {
	UserId   int `json:"-" gorm:"user_id"`
	FoodId   int `json:"-" gorm:"food_id"`
	Quantity int `json:"quantity" gorm:"quantity" binding:"required"`
}

func (c *CartCreate) TableName() string {
	return EntityName
}

type CartUpdate struct {
	UserId   int `json:"-" gorm:"user_id"`
	FoodId   int `json:"-" gorm:"food_id"`
	Quantity int `json:"quantity" gorm:"quantity" binding:"required"`
}

type CartChangeQuantity struct {
	UserId int `json:"-" gorm:"user_id"`
	FoodId int `json:"-" gorm:"food_id"`
}

func (c *CartChangeQuantity) TableName() string {
	return EntityName
}

func (c *CartUpdate) TableName() string {
	return EntityName
}

type CartDelete struct {
	UserId int `json:"-" gorm:"user_id"`
	FoodId int `json:"-" gorm:"food_id"`
}

func (c *CartDelete) TableName() string {
	return EntityName
}

var (
	ErrCartNotFound = common.NewCustomError(errors.New("cart not found"), "cart not found", "ErrCartNotFound")
)
