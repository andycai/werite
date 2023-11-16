package user

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return core.Render(c, "admin/login", fiber.Map{})
}
