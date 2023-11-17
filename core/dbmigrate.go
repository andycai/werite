package core

import (
	mc "github.com/andycai/werite/components/comment/model"
	mp "github.com/andycai/werite/components/page/model"
	ma "github.com/andycai/werite/components/post/model"
	mt "github.com/andycai/werite/components/tag/model"
	mu "github.com/andycai/werite/components/user/model"
	"gorm.io/gorm"
)

func AutoMigrate(dbs []*gorm.DB) {
	for _, db := range dbs {
		db.AutoMigrate(
			&mu.User{},
			&ma.Post{},
			&ma.Category{},
			&ma.PostTag{},
			&mp.Page{},
			&mt.Tag{},
			&mc.Comment{},
		)
	}
}
