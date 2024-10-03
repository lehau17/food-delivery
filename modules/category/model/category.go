package categorymodel

import (
	"errors"

	"github.com/lehau17/food_delivery/common"
)

const EntityName = "categories"

type Category struct {
	common.SqlModel

	Id          int           `json:"id" gorm:"column:id"`
	Name        string        `json:"name" gorm:"column:name"`
	Description string        `json:"description" gorm:"column:description"`
	Icon        *common.Image `json:"icon" gorm:"column:icon"`
}

func (c *Category) TableName() string {
	return EntityName
}

type CategoryCreate struct {
	Name        *string       `json:"name,omitempty" gorm:"column:name"`
	Description *string       `json:"description,omitempty" gorm:"column:description"`
	Icon        *common.Image `json:"icon,omitempty" gorm:"column:icon"`
}

type CategoryUpdate struct {
	Name        *string       `json:"name,omitempty" gorm:"column:name"`
	Description *string       `json:"description,omitempty" gorm:"column:description"`
	Icon        *common.Image `json:"icon,omitempty" gorm:"column:icon"`
}

func (c *CategoryCreate) TableName() string {
	return "categories"
}

func (c *CategoryUpdate) TableName() string {
	return "categories"
}

var (
	ErrCateNotFound = common.NewErrorResponse(errors.New("Category not found"), "Category not found", "Category not found", "ErrCateNotFound")
)
