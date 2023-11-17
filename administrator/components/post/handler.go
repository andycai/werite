package post

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
)

func PostsPage(c *fiber.Ctx) error {
	return core.Render(c, "admin/posts/posts", fiber.Map{
		//
	})
}

func PostPage(c *fiber.Ctx) error {
	return core.Render(c, "admin/posts/post", fiber.Map{
		//
	})
}
