package categorystorage

import (
	"context"

	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

func (s *sqlStore) FindCategory(ctx context.Context, conditions map[string]interface{}) (*categorymodel.Category, error) {
	db := s.db.Table(categorymodel.EntityName)
	var c categorymodel.Category
	if err := db.Where(conditions).Find(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
