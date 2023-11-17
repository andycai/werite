package model

const TableNamePostTag = "post_tag"

type PostTag struct {
	PostID int32 `gorm:"column:post_id;not null" json:"post_id"`
	TagID  int32 `gorm:"column:tag_id;not null" json:"tag_id"`
}

func (*PostTag) TableName() string {
	return TableNamePostTag
}
