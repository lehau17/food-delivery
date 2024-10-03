package categorystorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	db := s.db.Table(data.TableName())
	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
