package foodrepo

import (
	"context"

	foodmodel "github.com/lehau17/food_delivery/modules/food/model"
)

type FoodUpdateStore interface {
	UpdateFood(ctx context.Context, data *foodmodel.FoodUpdate, id int, user_id int) error
}

type FoodUpdateRepo struct {
	Repo FoodUpdateStore
}

func NewFoodUpdateRepo(repo FoodUpdateStore) *FoodUpdateRepo {
	return &FoodUpdateRepo{Repo: repo}
}
func (r *FoodUpdateRepo) UpdateFood(ctx context.Context, data *foodmodel.FoodUpdate, id int, user_id int) error {

	return r.Repo.UpdateFood(ctx, data, id, user_id)
}
