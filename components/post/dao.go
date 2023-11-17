package post

import "github.com/andycai/werite/components/post/model"

type PostDao struct{}

var Dao = new(PostDao)

func (ad PostDao) GetBySlug(slug string) (*model.Post, error) {
	var post model.Post
	err := db.Model(&post).
		Where("slug = ?", slug).
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
