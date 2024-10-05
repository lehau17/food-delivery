package foodbiz

import "context"

type FoodDeleteRepo interface {
	DeleteFood(ctx context.Context, id int, userId int) error
}

type FoodDeteteBiz struct {
	Repo FoodDeleteRepo
}

func NewFoodDeleteRepo(repo FoodDeleteRepo) *FoodDeteteBiz {
	return &FoodDeteteBiz{Repo: repo}
}

func (b *FoodDeteteBiz) DeleteFood(ctx context.Context, id int, userId int) error {
	return b.Repo.DeleteFood(ctx, id, userId)
}
