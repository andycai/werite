package system

import "github.com/andycai/werite/v2/model"

type PageSystem struct {
}

var Page = new(PageSystem)

func (ps PageSystem) GetBySlug(slug string) (*model.Page, error) {
	var page model.Page
	err := db.Model(&page).
		Where("slug = ?", slug).
		Find(&page).Error

	return &page, err
}
