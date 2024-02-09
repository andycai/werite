package core

import (
	"github.com/andycai/werite/model"
	"gorm.io/gorm"
)

func AutoMigrate(dbs []*gorm.DB) {
	for _, db := range dbs {
		db.AutoMigrate(
			&model.User{},
			// &model.Site{},
			&model.Post{},
			&model.Page{},
			&model.Media{},
			&model.Category{},
			&model.PostTag{},
			&model.Tag{},
			&model.Comment{},
		)
	}
}
