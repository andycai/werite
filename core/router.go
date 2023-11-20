package core

import (
	"github.com/andycai/werite/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	group := app.Group("/")
	for _, f := range routerNoCheckMap {
		f(group)
	}

	app.Use(middlewares.Authorize)
	for _, f := range routerCheckMap {
		f(group)
	}

	app.Use(middlewares.NotFoundRoute)
}
