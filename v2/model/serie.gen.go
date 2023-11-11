// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSerie = "serie"

// Serie mapped from table <serie>
type Serie struct {
	ID        int32     `gorm:"column:id;primaryKey" json:"id"`
	Slug      string    `gorm:"column:slug;not null" json:"slug"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Desc      string    `gorm:"column:desc;not null" json:"desc"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
}

// TableName Serie's table name
func (*Serie) TableName() string {
	return TableNameSerie
}