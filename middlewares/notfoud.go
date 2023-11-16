package middlewares

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": true,
		"msg":   "page not found",
	})
}
