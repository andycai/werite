// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameArticleTag = "article_tag"

// ArticleTag mapped from table <article_tag>
type ArticleTag struct {
	ArticleID int32 `gorm:"column:article_id;not null" json:"article_id"`
	TagID     int32 `gorm:"column:tag_id;not null" json:"tag_id"`
}

// TableName ArticleTag's table name
func (*ArticleTag) TableName() string {
	return TableNameArticleTag
}