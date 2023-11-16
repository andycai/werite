package model

import (
	"time"
)

const TableNameUser = "user"

type User struct {
	ID       int32     `gorm:"column:id;primaryKey" json:"id"`
	Username string    `gorm:"column:username;not null" json:"username"`
	Password string    `gorm:"column:password;not null;default:123456" json:"password"`
	Token    string    `gorm:"column:token;not null" json:"token"`
	Avatar   string    `gorm:"column:avatar;not null" json:"avatar"`
	Gender   int32     `gorm:"column:gender;not null;default:1" json:"gender"`
	Phone    string    `gorm:"column:phone;not null" json:"phone"`
	Email    string    `gorm:"column:email;not null" json:"email"`
	Addr     string    `gorm:"column:addr;not null" json:"addr"`
	IP       string    `gorm:"column:ip;not null" json:"ip"`
	LoginAt  time.Time `gorm:"column:login_at" json:"login_at"`
	LogoutAt time.Time `gorm:"column:logout_at" json:"logout_at"`
	CreateAt time.Time `gorm:"column:create_at;not null;default:current_timestamp" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;not null;default:current_timestamp" json:"update_at"`
	DeleteAt time.Time `gorm:"column:delete_at" json:"delete_at"`
	Name     string    `gorm:"column:name;not null" json:"name"`
}

func (*User) TableName() string {
	return TableNameUser
}
