package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameComment = "comment"

type Comment struct {
	ID        int32          `gorm:"column:id;primaryKey" json:"id"`
	UserID    int32          `gorm:"column:user_id;not null" json:"user_id"`
	Body      string         `gorm:"column:body;not null" json:"body"`
	Email     string         `gorm:"column:email;not null" json:"email"`
	IP        string         `gorm:"column:ip;not null" json:"ip"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:current_timestamp" json:"updated_at"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
}

func (*Comment) TableName() string {
	return TableNameComment
}
