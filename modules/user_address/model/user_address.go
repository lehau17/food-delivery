package useraddressmodel

import (
	"github.com/lehau17/food_delivery/common"
	citymodel "github.com/lehau17/food_delivery/modules/city/model"
)

const (
	EntityName = "user_addresses"
)

type UserAddress struct {
	common.SqlModel
	UserId int             `json:"-" gorm:"column:user_id"`
	User   *common.User    `json:"user" gorm:"preload:false"`
	CityId int             `json:"city_id" gorm:"column:city_id"`
	City   *citymodel.City `json:"city" gorm:"column:city"`
	Title  string          `json:"title" gorm:"column:title"`
	Icon   *common.Image   `json:"icon" gorm:"column:icon;type:json"`
	Addr   string          `json:"addr" gorm:"column:addr"`
	Lat    float64         `json:"lat" gorm:"column:lat"`
	Lng    float64         `json:"lng" gorm:"column:lng"`
}

func (ua *UserAddress) TableName() string {
	return EntityName
}
