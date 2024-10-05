package foodstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

func (s *sqlStore) UpdateFood(ctx context.Context, data *foodmodel.FoodUpdate, id int) error {
	db := s.db.Table(data.TableName())

	var foundFood foodmodel.Food
	if err := db.Where("id = ? AND status = 1", id).First(&foundFood).Error; err != nil {
		return common.ErrDb(err)
	}

	// Kiểm tra và gán các trường từ data vào foundFood
	if data.Name != nil {
		foundFood.Name = *data.Name
	}
	if data.Description != nil {
		foundFood.Description = *data.Description
	}
	if data.Price != nil {
		foundFood.Price = *data.Price
	}
	if data.Images != nil {
		foundFood.Images = data.Images
	}
	if data.CategoryId != nil {
		foundFood.CategoryId = *data.CategoryId
	}

	// Cập nhật vào cơ sở dữ liệu
	if err := db.Save(&foundFood).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
