package core

import (
	ma "github.com/andycai/werite/components/article/model"
	mc "github.com/andycai/werite/components/comment/model"
	mp "github.com/andycai/werite/components/page/model"
	mt "github.com/andycai/werite/components/tag/model"
	mu "github.com/andycai/werite/components/user/model"
	"gorm.io/gorm"
)

func AutoMigrate(dbs []*gorm.DB) {
	for _, db := range dbs {
		db.AutoMigrate(
			&mu.User{},
			&ma.Article{},
			&ma.ArticleTag{},
			&mp.Page{},
			&mt.Tag{},
			&mc.Comment{},
			&ma.Category{},
		)
	}
}
