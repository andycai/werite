package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTag = "tags"

type Tag struct {
	ID        uint           `gorm:"column:id;primaryKey" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	Posts     []Post         `gorm:"many2many:post_tag"`
}

func (*Tag) TableName() string {
	return TableNameTag
}
