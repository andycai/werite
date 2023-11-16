package model

const TableNameArticleTag = "article_tag"

type ArticleTag struct {
	ArticleID int32 `gorm:"column:article_id;not null" json:"article_id"`
	TagID     int32 `gorm:"column:tag_id;not null" json:"tag_id"`
}

func (*ArticleTag) TableName() string {
	return TableNameArticleTag
}
