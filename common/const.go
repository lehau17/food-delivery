package common

const (
	DB_RESTAURANT_TYPE = 1
	DB_USER_TYPE       = 2
	DB_CATEGORY_TYPE   = 3
	DB_FOOD_TYPE       = 4
)

const (
	CurrentUser               = "user"
	TopicUserLikeRestaurant   = "user-like-restaurant"
	TopicUserUnLikeRestaurant = "user-unlike-restaur"
	TopicUserLikeFood         = "user-like-food"
	TopicUserUnLikeFood       = "user-unlike-food"
)

type Requester interface {
	GetUId() int
	GetEmail() string
	GetRole() string
}
