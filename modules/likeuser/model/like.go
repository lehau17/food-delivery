package userlikerestaurantmodel

import (
	"errors"
	"time"

	"github.com/lehau17/food_delivery/common"
)

type Like struct {
	RestaurantId int          `json:"restaurant_id" gorm:"column:restaurant_id"`
	CreatedAt    *time.Time   `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UserId       int          `json:"user_id" gorm:"column:user_id"`
	User         *common.User `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

var (
	ErrLikeExist    = common.NewCustomError(errors.New("Like exists"), "Like exists", "ErrLikeExists")
	ErrLikeNotExist = common.NewCustomError(errors.New("Like not exists"), "Like not exists", "ErrLikeNotExists")
)

func (l *Like) GetResId() int {
	return l.RestaurantId
}
func (l *Like) GetUserId() int {
	return l.UserId
}
