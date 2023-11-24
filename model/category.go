package model

import (
	"gorm.io/gorm"
)

const TableNameCategory = "categories"

type Category struct {
	gorm.Model
	Slug string `gorm:"column:slug;not null" json:"slug"`
	Name string `gorm:"column:name;not null" json:"name"`
	Desc string `gorm:"column:desc;not null" json:"desc"`
}

func (*Category) TableName() string {
	return TableNameCategory
}
