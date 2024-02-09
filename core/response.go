package core

import (
	"github.com/gofiber/fiber/v2"
)

// func Err(c *Ctx, statusCode, code int) error {
// 	return c.Status(statusCode).JSON(fiber.Map{
// 		"code":  code,
// 		"error": enum.CodeText(code),
// 	})
// }

func Error(c *Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": err.Error(),
	})
}
