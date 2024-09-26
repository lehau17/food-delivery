package likestorage

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	userlikerestaurantmodel "github.com/lehau17/food_delivery/modules/likeuser/model"
)

func (s *sqlStore) CreateLike(ctx context.Context, data *userlikerestaurantmodel.Like) error {
	db := s.db
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
