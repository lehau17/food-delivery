package foodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) UpdateAvgPoint(ctx context.Context) error {
	db := s.db.Table(foodmodel.EntityName)
	if err := db.Exec(`UPDATE foods SET avg_point = (SELECT AVG(point) FROM food_ratings WHERE food_ratings.food_id = foods.id`).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil

}

func (s *sqlStore) UpdateAvgPointByFoodId(ctx context.Context, id int) error {
	db := s.db.Table(foodmodel.EntityName)

	// Thực hiện truy vấn SQL để cập nhật avg_point cho food có id cụ thể
	query := `
		UPDATE foods 
		SET avg_point = (
			SELECT AVG(point) 
			FROM food_ratings 
			WHERE food_ratings.food_id = ?
		)
		WHERE id = ?
	`

	if err := db.Exec(query, id, id).Error; err != nil {
		// Trả về lỗi từ database thông qua hàm common.ErrDb(err)
		return common.ErrDb(err)
	}

	return nil
}
