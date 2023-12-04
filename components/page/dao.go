package page

import (
	"fmt"

	"github.com/andycai/werite/model"
)

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

func (pd PageDao) GetAllByPage(page, numPerPage int) []model.Page {
	return pd.GetListByPage(page, numPerPage, "")
}

func (pd PageDao) GetListByPage(page, numPerPage int, q string) []model.Page {
	var pages []model.Page
	db.Model(&pages).
		Preload("User").
		Limit(numPerPage).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", q)).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&pages)

	return pages
}

func (pd PageDao) GetTrashListByPage(page, numPerPage int, q string) []model.Page {
	var pages []model.Page
	db.Model(&pages).
		Preload("User").
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", q)).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&pages)

	return pages
}

func (pd PageDao) DeleteByIds(ids []uint) {
	var page model.Page
	// pages := make([]model.Page, len(ids))
	// for i, v := range ids {
	// 	pages[i].ID = v
	// }
	db.Where(ids).Delete(&page)
}
