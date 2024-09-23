package usermodel

import "github.com/lehau17/food_delivery/common"

const EntityName = "users"

type User struct {
	common.SqlModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"-" gorm:"column:password"`
	Salt      string        `json:"-" gorm:"column:salt"`
	LastName  string        `json:"last_Name" gorm:"column:last_name"`
	FirstName string        `json:"first_name" gorm:"column:first_name"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Role      string        `json:"role" gorm:"column:role"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUid(common.DB_USER_TYPE)
}

type UserCreate struct {
	common.SqlModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"password" gorm:"column:password"`
	Salt      string        `json:"-" gorm:"column:salt"`
	LastName  string        `json:"last_Name" gorm:"column:last_name"`
	FirstName string        `json:"first_name" gorm:"column:first_name"`
	Role      string        `json:"-" gorm:"column:role"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUid(common.DB_USER_TYPE)
}
