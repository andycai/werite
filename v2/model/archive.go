package model

import "time"

type Archive struct {
	Time time.Time `gorm:"column:time;not null" bson:"time"`
}
