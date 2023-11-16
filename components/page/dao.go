package page

import "github.com/andycai/werite/v2/model"

type PageDao struct {
}

var Dao = new(PageDao)

func (pd PageDao) GetBySlug(slug string) (*model.Page, error) {
	var page model.Page
	err := db.Model(&page).
		Where("slug = ?", slug).
		Find(&page).Error

	return &page, err
}
