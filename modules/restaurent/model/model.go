package restaurentmodel

import "github.com/lehau17/food_delivery/common"

type Restaurant struct {
	common.SqlModel
	Name             string         `json:"name" gorm:"column:name;"`
	Addr             string         `json:"addr" gorm:"column:addr;"`
	CityId           int            `json:"city_id" gorm:"column:city_id;"` // tag
	Iat              int            `json:"iat" gorm:"iat"`
	Code             string         `json:"code" gorm:"column:code;"`
	Img              string         `json:"img" gorm:"column:img;"`
	Cover            *common.Images `json:"cover" gorm:"column:cover;"`
	Logo             *common.Image  `json:"logo" gorm:"column:logo;"`
	ShippingFeePerKm int            `json:"shipping_fee_per_km" gorm:"shipping_fee_per_km"`
	Price            float64        `json:"price" gorm:"column:price;"`
}

func (r *Restaurant) Mask(isAdminorOwner bool) {
	r.GenUid(common.DB_RESTAURANT_TYPE)
}

func (r *RestaurantCreate) Mask(isAdminorOwner bool) {
	r.GenUid(common.DB_RESTAURANT_TYPE)
}

type RestaurantCreate struct {
	common.SqlModel
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"addr" gorm:"column:addr;"`
	Code  string         `json:"code" gorm:"column:code;"`
	Price float64        `json:"price" gorm:"column:price;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
