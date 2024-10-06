package likefoodrepo

import (
	"context"

	"github.com/lehau17/food_delivery/common"
)

type UserStore interface {
	GetListUser(ctx context.Context, condition map[string]interface{}) ([]common.User, error)
}

type LikeFoodGetListUserStore interface {
	GetListUserLikeFood(ctx context.Context, foodId int) ([]int, error)
}

type LikeFoodGetListUserRepo struct {
	userStore     UserStore
	likeFoodStore LikeFoodGetListUserStore
}

func NewLikeFoodGetListUserRepo(userStore UserStore, likeFoodStore LikeFoodGetListUserStore) *LikeFoodGetListUserRepo {
	return &LikeFoodGetListUserRepo{userStore: userStore, likeFoodStore: likeFoodStore}
}

func (r *LikeFoodGetListUserRepo) GetListUserLikeFood(ctx context.Context, foodId int) ([]common.User, error) {
	ids, err := r.likeFoodStore.GetListUserLikeFood(ctx, foodId)
	if err != nil {
		return nil, err
	}
	if len(ids) <= 0 {
		return nil, nil
	}
	users, err := r.userStore.GetListUser(ctx, map[string]interface{}{"id": ids})
	if err != nil {
		return nil, err
	}
	return users, nil
}
