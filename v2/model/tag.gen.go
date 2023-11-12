// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTag = "tag"

// Tag mapped from table <tag>
type Tag struct {
	ID        int32          `gorm:"column:id;primaryKey" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
