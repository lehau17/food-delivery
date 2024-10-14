package usermodel

import (
	"errors"

	"github.com/lehau17/food_delivery/common"
)

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

func (u *User) GetUId() int {
	return u.Id
}
func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUid(common.DB_USER_TYPE)
}

type UserCreate struct {
	common.SqlModel
	Email     string        `json:"email" gorm:"column:email" binding:"required,email"` // Require email format
	Password  string        `json:"password" gorm:"column:password" binding:"required"` // Require password
	Salt      string        `json:"-" gorm:"column:salt"`
	LastName  string        `json:"last_name" gorm:"column:last_name" binding:"required"`   // Require last name
	FirstName string        `json:"first_name" gorm:"column:first_name" binding:"required"` // Require first name
	Role      string        `json:"-" gorm:"column:role"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

type UserVerifyOtp struct {
	Email string `json:"email" gorm:"column:email" binding:"required,email"` // Require email format
	Otp   string `json:"otp" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email" binding:"required,email"` // Require email format
	Password string `json:"password" gorm:"column:password" binding:"required"` // Require password
}

func (UserLogin) TableName() string {
	return "users"
}

func (UserCreate) TableName() string {
	return "users"
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUid(common.DB_USER_TYPE)
}

var (
	ErrUserExists    = common.NewCustomError(errors.New("User already exists"), "User already exists", "ErrUserExists")
	ErrUserNotExists = common.NewCustomError(errors.New("User not exists"), "User not exists", "ErrUserNotExists")
	ErrUserDisable   = common.NewCustomError(errors.New("User already disable"), "User already disable", "ErrUserDisable")
	ErrUserEnable    = common.NewCustomError(errors.New("User already enable"), "User already enable", "ErrUserEnable")
	ErrUserLoginFail = common.NewCustomError(errors.New("email or password incorrect"), "Email or password incorrect", "ErrLoginFail")
	ErrUserNotFound  = common.NewCustomError(errors.New("user not found"), "user not found", "ErrLoginFail")
	ErrOtp           = common.NewCustomError(errors.New("otp invalid"), "otp invalid", "ErrOtpInvalid")
)
