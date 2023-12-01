package page

import "github.com/andycai/werite/model"

type PageDao struct {
}

var Dao = new(PageDao)

func (ad PageDao) Count() int64 {
	var page model.Page
	var count int64
	db.Model(&page).Count(&count)

	return count
}

func (ad PageDao) CountTrash() int64 {
	var page model.Page
	var count int64
	db.Model(&page).Unscoped().Where("deleted_at IS NOT NULL").Count(&count)

	return count
}

func (pd PageDao) GetBySlug(slug string) (*model.Page, error) {
	var page model.Page
	err := db.Model(&page).
		Where("slug = ?", slug).
		Find(&page).Error

	return &page, err
}

func (pd PageDao) GetByID(id uint) (*model.Page, error) {
	var page model.Page
	err := db.Model(&page).
		Where("id = ?", id).
		Find(&page).Error

	return &page, err
}

func (pd PageDao) GetListByPage(page, numPerPage int) []model.Page {
	var pages []model.Page
	db.Model(&pages).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&pages)

	return pages
}

func (pd PageDao) GetTrashListByPage(page, numPerPage int) []model.Page {
	var pages []model.Page
	db.Model(&pages).
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&pages)

	return pages
}

func (pd PageDao) DeleteByIds(ids []uint) {
	pages := make([]model.Page, len(ids))
	for i, v := range ids {
		pages[i].ID = v
	}
	db.Delete(&pages)
}
