package ordermodel

import "github.com/lehau17/food_delivery/common"

const (
	EntityName = "orders"
)

type Order struct {
	common.SqlModel
	UserId     int          `json:"user_id" gorm:"column:user_id"`
	User       *common.User `json:"user" gorm:"-"`
	TotalPrice int          `json:"total_price" gorm:"column:total_price"`
	ShipperId  int          `json:"shipper_id" gorm:"column:shipper"`
	// Status     int          `json:"status" gorm:"column:status"`
}

func (o *Order) TableName() string {
	return EntityName
}

type OrderCreate struct {
	Id         int `json:"-" gorm:"column:id"`
	UserId     int `json:"user_id" gorm:"column:user_id"`
	TotalPrice int `json:"total_price" gorm:"column:total_price"`
	ShipperId  int `json:"shipper_id" gorm:"column:shipper"`
}

func (o *OrderCreate) TableName() string {
	return EntityName
}

type OrderUpdate struct {
	Id         int `json:"-" gorm:"column:id"`
	TotalPrice int `json:"total_price" gorm:"column:total_price"`
	ShipperId  int `json:"shipper_id" gorm:"column:shipper"`
}

func (o *OrderUpdate) TableName() string {
	return EntityName
}
