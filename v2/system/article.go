package system

import "github.com/andycai/werite/v2/model"

type ArticleSystem struct {
}

var Article = new(ArticleSystem)

func (us ArticleSystem) GetBySlug(slug string) (*model.Article, error) {
	var article model.Article
	err := db.Model(&article).
		Where("slug = ?", slug).
		Find(&article).Error

	return &article, err
}

func (as ArticleSystem) Count() int64 {
	// db := database.Get()
	// db.Debug().Model(&articles).
	// 	// Preload("Tags", func(db *gorm.DB) *gorm.DB {
	// 	// return db.Order("tags.name asc")
	// 	// }).
	// 	Limit(numPerPage).
	// 	Offset(page * numPerPage).
	// 	Order("created_at desc").
	// 	Find(&articles)

	var article model.Article
	var count int64
	db.Model(&article).Count(&count)

	return count
}

func (as ArticleSystem) GetListByPage(page, numPerPage int) []model.Article {
	var articles []model.Article
	db.Debug().Model(&articles).
		// Preload("Tags", func(db *gorm.DB) *gorm.DB {
		// return db.Order("tags.name asc")
		// }).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&articles)

	return articles
}
