package likestorage

import (
	"context"
	"log"

	"github.com/lehau17/food_delivery/common"
	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

func (s *sqlStore) GetRestautantsLike(context context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)
	db := s.db.Table(userlikerestaurantmodel.Like{}.TableName())
	type Data struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		Count        int `gorm:"column:count"`
	}
	var sqlData []Data
	if err := db.Select("restaurant_id, count(restaurant_id) as count").Where("restaurant_id in (?)", ids).Group("restaurant_id").Find(&sqlData).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for _, item := range sqlData {
		result[item.RestaurantId] = item.Count
	}
	log.Println(result)
	return result, nil

}
