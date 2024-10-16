package citymodel

import "github.com/lehau17/food_delivery/common"

type City struct {
	common.SqlModel
	Title string `json:"title" gorm:"title"`
}
