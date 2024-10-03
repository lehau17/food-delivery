package categortbiz

import (
	"context"

	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

type CategortUpdateStore interface {
	UpdateCate(ctx context.Context, data *categorymodel.CategoryUpdate, cateId int) error
}

type CategoryBizUpdate struct {
	store CategortUpdateStore
}

func NewCategoryBizUpdate(store CategortUpdateStore) *CategoryBizUpdate {
	return &CategoryBizUpdate{store: store}
}

func (c *CategoryBizUpdate) UpdateCategory(ctx context.Context, data *categorymodel.CategoryUpdate, cateId int) error {
	return c.store.UpdateCate(ctx, data, cateId)
}
