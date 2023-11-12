package router

import (
	"github.com/andycai/werite/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerCheck = append(routerCheck, registerHTMXCheckRouter)
	routerNoCheck = append(routerNoCheck, registerHTMXNoCheckRouter)
}

func registerHTMXNoCheckRouter(r fiber.Router) {
	r.Get("/htmx/home/tag-list", handler.HTMX.HomeTagList)
	r.Get("/htmx/home/global-feed", handler.HTMX.HomeGlobalFeed)
}

func registerHTMXCheckRouter(r fiber.Router) {
}
