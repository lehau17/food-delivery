package categorystorage

import (
	"context"
	"log"

	"github.com/lehau17/food_delivery/common"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

func (s *sqlStore) GetList(ctx context.Context, filter *categorymodel.Filter, paging *common.PagingCursor) ([]categorymodel.Category, error) {
	db := s.db.Table(categorymodel.EntityName)
	log.Println(filter)
	if f := filter; f != nil {
		if f.Name != "" {
			db = db.Where("name like ?", "%"+f.Name+"%")
		}
	}
	if paging.RealCursor != 0 {
		db = db.Where("id > ?", paging.RealCursor)
	}
	var data []categorymodel.Category
	if err := db.Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return data, nil
}
