package handler

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/gofiber/fiber/v2"
)

type Ctx = fiber.Ctx

// IP get remote IP
func IP(c *Ctx) string {
	return c.IP()
}

func Str(c *Ctx, key string, defaultValue ...string) string {
	return c.Params(key, defaultValue...)
}

func Int(c *Ctx, key string, defaultValue ...string) int {
	return cast.ToInt(c.Params(key, defaultValue...))
}

func Uint(c *Ctx, key string, defaultValue ...string) uint {
	return cast.ToUint(c.Params(key, defaultValue...))
}

func U32(c *Ctx, key string, defaultValue ...string) uint32 {
	return cast.ToUint32(c.Params(key, defaultValue...))
}

func I32(c *Ctx, key string, defaultValue ...string) int32 {
	return cast.ToInt32(c.Params(key, defaultValue...))
}

func U64(c *Ctx, key string, defaultValue ...string) uint64 {
	return cast.ToUint64(c.Params(key, defaultValue...))
}

func I64(c *Ctx, key string, defaultValue ...string) int64 {
	return cast.ToInt64(c.Params(key, defaultValue...))
}

// Msg push common response
func Msg(c *Ctx, code int, msg string) error {
	return c.JSON(fiber.Map{
		"code": code,
		"msg":  msg,
	})
}

func render(c *Ctx, name string, bind interface{}, layouts ...string) error {
	return c.Render(fmt.Sprintf("%s", name), bind, layouts...)
}
