package restaurentstorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

func (s *sqlStore) GetList(context context.Context, filter *restaurentmodel.Filter, paging *common.Paging, more ...string) (result []restaurentmodel.Restaurant, err error) {
	db := s.db.Table(restaurentmodel.Restaurant{}.TableName())
	for i := range more {
		db = db.Preload(more[i])
	}
	db = db.Where("status in (?)", filter.Status).Order("id desc")
	if f := filter; f != nil {
		// custom db

	}
	//get result and return
	err = db.Offset((paging.Page - 1) * paging.Limit).Find(&result).Limit(paging.Limit).Error
	if err != nil {
		return nil, err
	}

	//add total into paging
	if err = db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	return result, nil
}
