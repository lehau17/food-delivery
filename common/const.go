package common

const (
	DB_RESTAURANT_TYPE = 1
	DB_USER_TYPE       = 2
)

const (
	CurrentUser               = "user"
	TopicUserLikeRestaurant   = "user-like-restaurant"
	TopicUserUnLikeRestaurant = "user-unlike-restaur"
)

type Requester interface {
	GetUId() int
	GetEmail() string
	GetRole() string
}
