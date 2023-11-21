package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePage = "pages"

type Page struct {
	ID          uint           `gorm:"column:id;primaryKey" json:"id"`
	UserID      uint           `gorm:"column:user_id;not null" json:"user_id"`
	Slug        string         `gorm:"column:slug;not null" json:"slug"`
	Title       string         `gorm:"column:title;not null" json:"title"`
	Body        string         `gorm:"column:body;not null" json:"body"`
	PublishedAt time.Time      `gorm:"column:published_at" json:"published_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
}

func (*Page) TableName() string {
	return TableNamePage
}
