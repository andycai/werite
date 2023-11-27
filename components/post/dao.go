package post

import (
	"github.com/andycai/werite/model"
	"gorm.io/gorm"
)

type PostDao struct{}

var Dao = new(PostDao)

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

func (ad PostDao) Count() int64 {
	// db := database.Get()
	// db.Debug().Model(&posts).
	// 	// Preload("Tags", func(db *gorm.DB) *gorm.DB {
	// 	// return db.Order("tags.name asc")
	// 	// }).
	// 	Limit(numPerPage).
	// 	Offset(page * numPerPage).
	// 	Order("created_at desc").
	// 	Find(&posts)

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
