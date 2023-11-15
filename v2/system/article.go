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
