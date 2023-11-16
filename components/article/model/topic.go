package model

import (
	"time"
)

const TableNameTopic = "topic"

type Topic struct {
	ID        int32     `gorm:"column:id;primaryKey" json:"id"`
	Slug      string    `gorm:"column:slug;not null" json:"slug"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Desc      string    `gorm:"column:desc;not null" json:"desc"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
}

func (*Topic) TableName() string {
	return TableNameTopic
}
