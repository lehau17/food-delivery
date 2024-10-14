package common

const (
	DB_RESTAURANT_TYPE        = 1
	DB_USER_TYPE              = 2
	DB_CATEGORY_TYPE          = 3
	DB_FOOD_TYPE              = 4
	DB_FOOD_RATING_TYPE       = 5
	DB_RESTAURANT_RATING_TYPE = 6
)

const (
	CurrentUser                   = "user"
	TopicUserLikeRestaurant       = "user-like-restaurant"
	TopicUserUnLikeRestaurant     = "user-unlike-restaur"
	TopicUserLikeFood             = "user-like-food"
	TopicUserUnLikeFood           = "user-unlike-food"
	TopicCreateFoodRestaurant     = "create-food-restaurant"
	TopicDeleteFoodRestaurant     = "delete-food-restaurant"
	TopicSendMailWhileUserCreated = "send-mail-while-user-created"
)

type Requester interface {
	GetUId() int
	GetEmail() string
	GetRole() string
}
