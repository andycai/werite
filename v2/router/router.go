package router

import (
	"github.com/andycai/werite/v2/middleware"
	"github.com/gofiber/fiber/v2"
)

var (
	routerNoCheck = make([]func(fiber.Router), 0)
	routerCheck   = make([]func(fiber.Router), 0)
)

func Setup(app *fiber.App) {
	v2 := app.Group("/")
	// 需要登录校验的路由
	for _, f := range routerCheck {
		f(v2)
	}

	// v2 := app.Group("/api", JWTAuthMiddleware)
	// 不需要登录校验的路由
	for _, f := range routerNoCheck {
		f(v2)
	}
	app.Use(middleware.NotFoundRoute)
}
