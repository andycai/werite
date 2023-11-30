package post

import (
	"github.com/andycai/werite/model"
	"gorm.io/gorm"
)

type PostDao struct{}

var Dao = new(PostDao)

//#region Post

func (ad PostDao) GetBySlug(slug string) (*model.Post, error) {
	var post model.Post
	err := db.Model(&post).
		Where("slug = ?", slug).
		Find(&post).Error

	return &post, err
}

func (ad PostDao) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	err := db.Model(&post).
		Where("id = ?", id).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Find(&post).Error

	return &post, err
}

func (ad PostDao) CountByPublished() int64 {
	return ad.countByDraft(0)
}

func (ad PostDao) CountByDraft() int64 {
	return ad.countByDraft(1)
}

func (ad PostDao) countByDraft(draft int) int64 {
	var post model.Post
	var count int64
	db.Model(&post).Where("is_draft = ?", draft).Count(&count)

	return count
}

func (ad PostDao) CountByTrash() int64 {
	var post model.Post
	var count int64
	db.Model(&post).Unscoped().Where("deleted_at IS NOT NULL").Count(&count)

	return count
}

func (ad PostDao) Count() int64 {
	var post model.Post
	var count int64
	db.Model(&post).Count(&count)

	return count
}

func (ad PostDao) GetListByPage(page, numPerPage int) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

func (ad PostDao) GetPublishedListByPage(page, numPerPage int) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Where("is_draft = ?", 0).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

func (ad PostDao) GetDraftListByPage(page, numPerPage int) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Where("is_draft = ?", 1).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

func (ad PostDao) GetTrashListByPage(page, numPerPage int) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

//#endregion

//#region Category

func (ad PostDao) CatgegoryCount() int64 {
	var category model.Category
	var count int64
	db.Model(&category).Count(&count)

	return count
}

func (ad PostDao) GetCategories() []model.Category {
	var categories []model.Category
	db.Model(&categories).
		Find(&categories)

	return categories
}

func (ad PostDao) GetCategoriesByPage(page, numPerPage int) []model.Category {
	var categories []model.Category
	db.Model(&categories).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Find(&categories)

	return categories
}

func (ad PostDao) GetCategoryByID(id uint) (*model.Category, error) {
	var categoryVo model.Category
	err := db.Model(&categoryVo).Where("id = ?", id).Find(&categoryVo).Error

	return &categoryVo, err
}

//#region Tag

func (ad PostDao) TagCount() int64 {
	var tag model.Tag
	var count int64
	db.Model(&tag).Count(&count)

	return count
}

func (ad PostDao) GetTagsByPage(page, numPerPage int) []model.Tag {
	var tags []model.Tag
	db.Model(&tags).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Find(&tags)

	return tags
}

func (ad PostDao) GetTagByID(id uint) (*model.Tag, error) {
	var tagVo model.Tag
	err := db.Model(&tagVo).Where("id = ?", id).Find(&tagVo).Error

	return &tagVo, err
}

//#endregion
