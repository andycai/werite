package media

import (
	"net/http"
	"path/filepath"

	"github.com/andycai/werite/administrator/components/user"
	"github.com/andycai/werite/conf"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/log"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
)

type SimpleMedia struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Url  string `json:"url"`
}

func handleQuery(c *fiber.Ctx) error {
	path := c.Query("path")
	medias, err := getLatest(path, conf.GetIntValue("app.mediaShowLimit", 20))
	if err != nil {
		return core.Error(c, http.StatusInternalServerError, err)
	}

	var result []SimpleMedia = make([]SimpleMedia, len(medias))
	for i, media := range medias {
		mediaHost := conf.GetValue("app.mediaHost", "http://127.0.0.1:8000")
		mediaPrefix := conf.GetValue("app.mediaPrefix", "/media/")
		media.BuildPublicUrls(mediaHost, mediaPrefix)

		result[i] = SimpleMedia{
			Name: media.Name,
			Size: media.Size,
			Url:  media.PublicUrl,
		}
	}

	return c.JSON(fiber.Map{
		"medias": result,
	})
}

func handleMedia(c *fiber.Ctx) error {
	fullPath := c.Params("*")
	path, name := filepath.Split(fullPath)
	if len(path) > 1 && path[0] != '/' {
		path = "/" + path
	}
	img, err := getMedia(path, name)
	if err != nil {
		return core.Error(c, http.StatusNotFound, err)
	}

	if img.External {
		return c.Redirect(img.StorePath)
	}

	uploadDir := KEY_CMS_UPLOAD_DIR
	filepath := filepath.Join(uploadDir, img.StorePath)
	return c.SendFile(filepath)
}

func handleUpload(c *fiber.Ctx) error {
	created := c.Query("created")
	path := c.Query("path")
	name := c.Query("name")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	mFile, err := file.Open()
	if err != nil {
		return err
	}
	defer mFile.Close()

	if path == "" {
		path = "/"
	}
	if name == "" {
		name = file.Filename
	}
	r, err := uploadFile(path, name, mFile)
	if err != nil {
		return err
	}

	var media model.Media

	userVo := user.Current(c)
	media.Name = r.Name
	media.Path = r.Path
	media.External = r.External
	media.StorePath = r.StorePath
	media.Size = r.Size
	media.ContentType = r.ContentType
	media.Dimensions = r.Dimensions
	media.Directory = false
	media.Ext = r.Ext
	media.ContentType = r.ContentType
	media.Published = true

	if userVo != nil {
		media.Creator = *userVo
		media.CreatorID = userVo.ID
	}

	if created != "" {
		result := db.Create(&media)
		if result.Error != nil {
			return result.Error
		}
	}

	mediaHost := conf.GetValue("app.mediaHost", "http://127.0.0.1:8000")
	mediaPrefix := conf.GetValue("app.mediaPrefix", "/media/")
	media.BuildPublicUrls(mediaHost, mediaPrefix)

	r.PublicUrl = media.PublicUrl
	r.Thumbnail = media.Thumbnail

	return c.JSON(fiber.Map{
		"title":   r.Name,
		"isImage": r.ContentType == "image",
		"url":     r.PublicUrl,
		"bytes":   r.Size,
	})
}

func handleDelete(c *fiber.Ctx) error {
	path := c.Query("path")
	name := c.Query("name")

	media, err := getMedia(path, name)
	if err != nil {
		return core.Error(c, http.StatusNotFound, err)
	}

	if err := removeFile(path, name); err != nil {
		log.Infof("Delete file failed: %s, %s", media.StorePath, err)
	}

	db.Where("name = ? AND path = ?", name, path).Delete(&media)

	return nil
}
