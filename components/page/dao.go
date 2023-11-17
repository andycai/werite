package page

import "github.com/andycai/werite/components/page/model"

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

func (pd PageDao) GetListByPage(page, numPerPage int) []model.Page {
	var pages []model.Page
	db.Debug().Model(&pages).
		// Preload("Tags", func(db *gorm.DB) *gorm.DB {
		// return db.Order("tags.name asc")
		// }).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&pages)

	return pages
}
