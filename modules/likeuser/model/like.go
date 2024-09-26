package userlikerestaurantmodel

import (
	"errors"
	"time"

	"github.com/lehau17/food_delivery/common"
)

type Like struct {
	RestaurantId int          `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int          `json:"user_id" gorm:"column:user_id"`
	CreatedAt    *time.Time   `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	User         *common.User `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

var (
	ErrLikeExist = common.NewCustomError(errors.New("Like exists"), "Like exists", "ErrLikeExists")
)
