package categorystorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

func (s *sqlStore) RemoveCategory(ctx context.Context, id int) error {
	db := s.db.Table(categorymodel.EntityName)
	if err := db.Where("id = ? and status = 1", id).Update("status", 0).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
