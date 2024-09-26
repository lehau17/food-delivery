package common

const (
	DB_RESTAURANT_TYPE = 1
	DB_USER_TYPE       = 2
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUId() int
	GetEmail() string
	GetRole() string
}
