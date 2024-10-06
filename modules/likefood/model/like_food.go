package likefoodmodel

import (
	"errors"
	"time"

	"github.com/lehau17/food_delivery/common"
)

const (
	EntityName = "food_likes"
)

type LikeFood struct {
	FoodId    int          `json:"food_id" gorm:"column:food_id"`
	UserId    int          `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time   `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time   `json:"updated_at" gorm:"column:updated_at;autoCreateTime"`
	User      *common.User `json:"user" gorm:"preload:false"`
}

func (lf *LikeFood) TableName() string {
	return EntityName
}

type LikeFoodCreate struct {
	FoodId int `json:"-" gorm:"column:food_id"`
	UserId int `json:"-" gorm:"column:user_id"`
}

func (lf *LikeFoodCreate) TableName() string {
	return EntityName
}

var (
	ErrLikeExists = common.NewErrorResponse(errors.New("like exists"), "Like Exists", "Like food already exists", "ErrLikeFoodExists")
)
