package categortbiz

import (
	"context"

	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

type CategortCreateStore interface {
	CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type CategoryBizCreate struct {
	store CategortCreateStore
}

func NewCategoryBizCreate(store CategortCreateStore) *CategoryBizCreate {
	return &CategoryBizCreate{store: store}
}

func (c *CategoryBizCreate) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	return c.store.CreateCategory(ctx, data)
}
