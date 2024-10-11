package restaurantratingmodel

import (
	"errors"

	"github.com/lehau17/food_delivery/common"
)

const (
	EntityName = "restaurant_ratings"
)

type RestaurantRating struct {
	common.SqlModel
	Point        float32      `json:"point" gorm:"column:point"`
	Comment      string       `json:"comment" gorm:"column:comment"`
	UserId       int          `json:"-" gorm:"column:user_id"`
	User         *common.User `json:"user" gorm:"preload:false"`
	RestaurantId int          `json:"-" gorm:"column:restaurant_id"`
	// User         *common.User `json:"user" gorm:"preload:false"`
}

func (r *RestaurantRating) TableName() string {
	return EntityName
}

func (r *RestaurantRating) Mask() {
	r.GenUid(common.DB_RESTAURANT_RATING_TYPE)
	r.User.Mask(false)
}

type RestaurantRatingCreate struct {
	Point        float32 `json:"point" gorm:"column:point"`
	Comment      string  `json:"comment" gorm:"column:comment"`
	UserId       int     `json:"-" gorm:"column:user_id"`
	RestaurantId int     `json:"-" gorm:"column:restaurant_id"`
}

func (r *RestaurantRatingCreate) TableName() string {
	return EntityName
}

var (
	ErrRestaurantRatingNotExist = common.NewFullErrorResponse(400, errors.New("restaurant rating not exist"), "restaurant rating not exist", "restaurant rating not exist", "ErrResRatingNotExist")
)
