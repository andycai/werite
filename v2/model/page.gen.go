// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePage = "page"

// Page mapped from table <page>
type Page struct {
	ID        int32          `gorm:"column:id;primaryKey" json:"id"`
	UserID    int32          `gorm:"column:user_id;not null" json:"user_id"`
	Slug      string         `gorm:"column:slug;not null" json:"slug"`
	Title     string         `gorm:"column:title;not null" json:"title"`
	Body      string         `gorm:"column:body;not null" json:"body"`
	IsDraft   int32          `gorm:"column:is_draft;not null" json:"is_draft"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;not null" json:"deleted_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
}

// TableName Page's table name
func (*Page) TableName() string {
	return TableNamePage
}
