package foodstorage

// import (
// 	"context"

// 	"github.com/lehau17/food_delivery/common"
// 	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
// 	likefoodmodel "github.com/lehau17/food_delivery/modules/likefood/model"
// )

// func (s *sqlStore) GetListFoodUserLiked(ctx context.Context, userId int) ([]foodmodel.Food, error) {
// 	db := s.db.Where(likefoodmodel.EntityName)
// 	var ids []int
// 	if err := db.Where("user_id = ?", userId).Select("food_id").Find(&ids).Error; err != nil {
// 		return nil, common.ErrDb(err)
// 	}
// 	return ids, nil
// }
