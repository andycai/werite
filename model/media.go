package model

import (
	"database/sql"
	"path/filepath"
	"time"
)

type Media struct {
	UpdatedAt   time.Time    `json:"updated_at" gorm:"index"`
	CreatedAt   time.Time    `json:"created_at" gorm:"index"`
	Thumbnail   string       `json:"thumbnail,omitempty" gorm:"size:500"`
	Tags        string       `json:"tags,omitempty" gorm:"size:200;index"`
	Title       string       `json:"title,omitempty" gorm:"size:200"`
	Alt         string       `json:"alt,omitempty"`
	Description string       `json:"description,omitempty"`
	Keywords    string       `json:"keywords,omitempty"`
	CreatorID   uint         `json:"-"`
	Creator     User         `json:"-"`
	Author      string       `json:"author" gorm:"size:64"`
	Published   bool         `json:"published"`
	PublishedAt sql.NullTime `json:"published_at"`
	ContentType string       `json:"content_type" gorm:"size:32"`
	Remark      string       `json:"remark"`
	Size        int64        `json:"size"`
	Directory   bool         `json:"directory" gorm:"index"`
	Path        string       `json:"path" gorm:"size:200;uniqueIndex:,composite:_path_name"`
	Name        string       `json:"name" gorm:"size:200;uniqueIndex:,composite:_path_name"`
	Ext         string       `json:"ext" gorm:"size:100"`
	Dimensions  string       `json:"dimensions" gorm:"size:200"` // x*y
	StorePath   string       `json:"-" gorm:"size:300"`
	External    bool         `json:"external"`
	PublicUrl   string       `json:"public_url,omitempty" gorm:"-"`
}
type MediaFolder struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	FilesCount   int64  `json:"filesCount"`
	FoldersCount int64  `json:"foldersCount"`
}

func (m *Media) BuildPublicUrls(mediaHost string, mediaPrefix string) {
	if m.Directory {
		m.PublicUrl = ""
		return
	}

	publicUrl := filepath.Join(mediaPrefix, m.Path, m.Name)
	if mediaHost != "" {
		if mediaHost[len(mediaHost)-1] == '/' {
			mediaHost = mediaHost[:len(mediaHost)-1]
		}
		publicUrl = mediaHost + publicUrl
	}
	m.PublicUrl = publicUrl

	if m.ContentType == "image" && m.Thumbnail == "" {
		m.Thumbnail = m.PublicUrl
	}
}

type UploadResult struct {
	PublicUrl   string `json:"publicUrl"`
	Thumbnail   string `json:"thumbnail"`
	Path        string `json:"path"`
	Name        string `json:"name"`
	External    bool   `json:"external"`
	StorePath   string `json:"storePath"`
	Dimensions  string `json:"dimensions"`
	Ext         string `json:"ext"`
	Size        int64  `json:"size"`
	ContentType string `json:"contentType"`
}
