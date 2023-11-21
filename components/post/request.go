package post

import (
	"time"

	"github.com/andycai/werite/components/post/model"
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type requestCreate struct {
	ID          uint   `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Body        string `json:"body" validate:"required"`
	IsDraft     uint   `json:"is_draft" form:"is_draft" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
	PublishedAt string `json:"published_at" form:"published_at" validate:"required"`
}

func Bind(c *fiber.Ctx, post *model.Post) error {
	var r requestCreate
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	if err := core.Validate(r); err != nil {
		return err
	}

	post.ID = r.ID
	post.Title = r.Title
	post.Description = r.Description
	post.Body = r.Body
	post.IsDraft = r.IsDraft
	post.CategoryID = r.CategoryID

	if r.Slug != "" {
		post.Slug = r.Slug
	} else {
		post.Slug = slug.Make(r.Title)
	}

	if r.PublishedAt != "" {
		post.PublishedAt = core.ParseDate(r.PublishedAt)
	} else {
		post.PublishedAt = time.Now()
	}

	return nil
}
