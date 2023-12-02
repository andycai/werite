package post

import (
	"fmt"

	"github.com/andycai/werite/model"
	"gorm.io/gorm"
)

type PostDao struct{}

var Dao = new(PostDao)

//#region Post

func (pd PostDao) GetBySlug(slug string) (*model.Post, error) {
	var post model.Post
	err := db.Model(&post).
		Where("slug = ?", slug).
		Find(&post).Error

	return &post, err
}

func (pd PostDao) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	err := db.Model(&post).
		Where("id = ?", id).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Find(&post).Error

	return &post, err
}

func (pd PostDao) CountByPublished() int64 {
	return pd.countByDraft(0)
}

func (pd PostDao) CountByDraft() int64 {
	return pd.countByDraft(1)
}

func (pd PostDao) countByDraft(draft int) int64 {
	var post model.Post
	var count int64
	db.Model(&post).Where("is_draft = ?", draft).Count(&count)

	return count
}

func (pd PostDao) CountByTrash() int64 {
	var post model.Post
	var count int64
	db.Model(&post).Unscoped().Where("deleted_at IS NOT NULL").Count(&count)

	return count
}

func (pd PostDao) Count() int64 {
	var post model.Post
	var count int64
	db.Model(&post).Count(&count)

	return count
}

func (pd PostDao) GetAllByPage(page, numPerPage int) []model.Post {
	return pd.GetListByPage(page, numPerPage, 0, "")
}

func (pd PostDao) GetListByPage(page, numPerPage int, category int, q string) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Where("category = ?", category).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", q)).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

func (pd PostDao) GetPublishedListByPage(page, numPerPage int, category int, q string) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Where("is_draft = ?", 0).
		Where("category = ?", category).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", q)).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

func (pd PostDao) GetDraftListByPage(page, numPerPage int, category int, q string) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Where("is_draft = ?", 1).
		Where("category = ?", category).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", q)).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

func (pd PostDao) GetTrashListByPage(page, numPerPage int, category int, q string) []model.Post {
	var posts []model.Post
	db.Model(&posts).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Where("category = ?", category).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", q)).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}

func (pd PostDao) DeleteByIds(ids []uint) {
	posts := make([]model.Post, len(ids))
	for i, v := range ids {
		posts[i].ID = v
	}
	db.Delete(&posts)
}

//#endregion

//#region Category

func (pd PostDao) CountCatgegory() int64 {
	var category model.Category
	var count int64
	db.Model(&category).Count(&count)

	return count
}

func (pd PostDao) GetCategories() []model.Category {
	var categories []model.Category
	db.Model(&categories).
		Find(&categories)

	return categories
}

func (pd PostDao) GetCategoriesByPage(page, numPerPage int) []model.Category {
	var categories []model.Category
	db.Model(&categories).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Find(&categories)

	return categories
}

func (pd PostDao) GetCategoryByID(id uint) (*model.Category, error) {
	var categoryVo model.Category
	err := db.Model(&categoryVo).Where("id = ?", id).Find(&categoryVo).Error

	return &categoryVo, err
}

func (pd PostDao) DeleteCategoriesByIds(ids []uint) {
	categories := make([]model.Category, len(ids))
	for i, v := range ids {
		categories[i].ID = v
	}
	db.Delete(&categories)
}

//#region Tag

func (pd PostDao) CountTag() int64 {
	var tag model.Tag
	var count int64
	db.Model(&tag).Count(&count)

	return count
}

func (pd PostDao) GetTagsByPage(page, numPerPage int) []model.Tag {
	var tags []model.Tag
	db.Model(&tags).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Find(&tags)

	return tags
}

func (pd PostDao) GetTagByID(id uint) (*model.Tag, error) {
	var tagVo model.Tag
	err := db.Model(&tagVo).Where("id = ?", id).Find(&tagVo).Error

	return &tagVo, err
}

func (pd PostDao) DeleteTagsByIds(ids []uint) {
	tags := make([]model.Tag, len(ids))
	for i, v := range ids {
		tags[i].ID = v
	}
	db.Delete(&tags)
}

//#endregion
