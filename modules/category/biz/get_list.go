package categortbiz

import (
	"context"

	"github.com/lehau17/food_delivery/common"
	categorymodel "github.com/lehau17/food_delivery/modules/category/model"
)

type CategoryGetListStore interface {
	GetList(ctx context.Context, filter *categorymodel.Filter, paging *common.PagingCursor) ([]categorymodel.Category, error)
}

type CategoryGetListBiz struct {
	store CategoryGetListStore
}

func NewCategoryGetListBiz(store CategoryGetListStore) *CategoryGetListBiz {
	return &CategoryGetListBiz{store: store}
}

func (c *CategoryGetListBiz) GetList(ctx context.Context, filter *categorymodel.Filter, paging *common.PagingCursor) ([]categorymodel.Category, error) {
	return c.store.GetList(ctx, filter, paging)
}
