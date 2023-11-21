package core

import (
	mpa "github.com/andycai/werite/components/page/model"
	mpo "github.com/andycai/werite/components/post/model"
	mu "github.com/andycai/werite/components/user/model"
	"gorm.io/gorm"
)

func AutoMigrate(dbs []*gorm.DB) {
	for _, db := range dbs {
		db.AutoMigrate(
			&mu.User{},
			&mpo.Post{},
			&mpo.Category{},
			&mpo.PostTag{},
			&mpa.Page{},
			&mpo.Tag{},
			&mpo.Comment{},
		)
	}
}
