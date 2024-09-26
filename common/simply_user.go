package common

type User struct {
	SqlModel
	LastName  string `json:"last_Name" gorm:"column:last_name"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	Role      string `json:"role" gorm:"column:role"`
	Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) Mask(isAdmin bool) {
	u.GenUid(DB_USER_TYPE)
}
