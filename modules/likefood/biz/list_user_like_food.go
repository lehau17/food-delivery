package likefoodbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
)

type LikeFoodListUserRepo interface {
	GetListUserLikeFood(ctx context.Context, foodId int) ([]common.User, error)
}

type LikeFoodListUserBiz struct {
	repo LikeFoodListUserRepo
}

func NewLikeFoodListUserBiz(repo LikeFoodListUserRepo) *LikeFoodListUserBiz {
	return &LikeFoodListUserBiz{repo: repo}
}

func (b *LikeFoodListUserBiz) GetListUserLikeFood(ctx context.Context, foodId int) ([]common.User, error) {
	return b.repo.GetListUserLikeFood(ctx, foodId)
}
