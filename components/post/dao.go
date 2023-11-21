package post

import (
	mpo "github.com/andycai/werite/components/post/model"
)

type PostDao struct{}

var Dao = new(PostDao)

func (ad PostDao) GetBySlug(slug string) (*mpo.Post, error) {
	var post mpo.Post
	err := db.Model(&post).
		Where("slug = ?", slug).
		Find(&post).Error

	return &post, err
}

func (ad PostDao) GetByID(id uint) (*mpo.Post, error) {
	var post mpo.Post
	err := db.Model(&post).
		Where("id = ?", id).
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

	var post mpo.Post
	var count int64
	db.Model(&post).Count(&count)

	return count
}

func (ad PostDao) GetListByPage(page, numPerPage int) []mpo.Post {
	var posts []mpo.Post
	db.Debug().Model(&posts).
		// Preload("Tags", func(db *gorm.DB) *gorm.DB {
		// return db.Order("tags.name asc")
		// }).
		Limit(numPerPage).
		Offset(page * numPerPage).
		Order("created_at desc").
		Find(&posts)

	return posts
}
