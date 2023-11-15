package system

import "gorm.io/gorm"

var db *gorm.DB

func SetDB(gormdb *gorm.DB) {
	db = gormdb
}
