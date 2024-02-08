package media

import (
	"github.com/andycai/werite/administrator/components/user"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
)

func ManagerPage(c *fiber.Ctx) error {
	return nil
}

func Upload(c *fiber.Ctx) error {
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

	mediaHost := ""          // user.GetValue(enum.KEY_CMS_MEDIA_HOST)
	mediaPrefix := "/media/" // user.GetValue(enum.KEY_CMS_MEDIA_PREFIX)
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

func Delete(c *fiber.Ctx) error {
	return nil
}