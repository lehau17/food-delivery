package categortbiz

import (
	"context"
	"log"

	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

type CategoryDeleteStore interface {
	RemoveCategory(ctx context.Context, id int) error
	FindCategory(ctx context.Context, conditions map[string]interface{}) (*categorymodel.Category, error)
}

type CategoryDeleteBiz struct {
	store CategoryDeleteStore
}

func NewCategoryDeleteBiz(store CategoryDeleteStore) *CategoryDeleteBiz {
	return &CategoryDeleteBiz{store: store}
}

func (b *CategoryDeleteBiz) DeleteCategory(ctx context.Context, id int) error {
	foundCate, err := b.store.FindCategory(ctx, map[string]interface{}{"status": 1, "id": 1})
	if err != nil {
		return nil
	}
	log.Println("Found category>>>>", foundCate)
	if foundCate.Id == 0 {
		return categorymodel.ErrCateNotFound
	}
	return b.store.RemoveCategory(ctx, id)
}
