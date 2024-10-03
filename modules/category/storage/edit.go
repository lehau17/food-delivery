package categorystorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

func (s *sqlStore) UpdateCate(ctx context.Context, data *categorymodel.CategoryUpdate, cateId int) error {
	db := s.db.Table(data.TableName())
	// get Cate
	var cate *categorymodel.Category
	if err := db.Where("id = ?", cateId).Find(&cate).Error; err != nil {
		return common.ErrDb(err)
	}

	// check Cate
	if cate == nil {
		return categorymodel.ErrCateNotFound
	}

	// update
	// Kiểm tra và cập nhật chỉ những trường có giá trị
	if data.Name != nil {
		cate.Name = *data.Name
	}
	if data.Description != nil {
		cate.Description = *data.Description
	}
	if data.Icon != nil {
		cate.Icon = data.Icon
	}

	// Lưu lại thay đổi
	if err := db.Save(&cate).Error; err != nil {
		return common.ErrDb(err)

	}
	return nil
}
