package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArticle = "article"

type Article struct {
	ID          int32          `gorm:"column:id;primaryKey" json:"id"`
	UserID      int32          `gorm:"column:user_id;not null" json:"user_id"`
	Slug        string         `gorm:"column:slug;not null" json:"slug"`
	Title       string         `gorm:"column:title;not null" json:"title"`
	Description string         `gorm:"column:description;not null" json:"description"`
	Body        string         `gorm:"column:body;not null" json:"body"`
	TopicID     int32          `gorm:"column:topic_id" json:"topic_id"`
	IsDraft     int32          `gorm:"column:is_draft" json:"is_draft"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
}

func (*Article) TableName() string {
	return TableNameArticle
}
