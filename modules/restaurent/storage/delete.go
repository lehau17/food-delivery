package restaurentstorage

import (
	"context"
	"errors"

	restaurentmodel "github.com/lehau17/food_delivery/modules/restaurent/model"
)

func (s *sqlStore) DeleteRestaurant(context context.Context, id int) error {
	db := s.db.Table(restaurentmodel.Restaurant{}.TableName())
	if id <= 0 {
		return errors.New("invalid request")
	}
	var data restaurentmodel.Restaurant
	//find existing
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return errors.New("invalid request")
	}

	if data.Id == 0 {
		return errors.New("user not found")

	}

	if data.Status == 0 {
		return errors.New("user not found")

	}

	// execute command
	if err := db.Table(restaurentmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("status", 0).Error; err != nil {
		return err
	}
	return nil
}
