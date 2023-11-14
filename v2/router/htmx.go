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
	r.Get("/htmx/sign-in", handler.HTMX.SignInPage)
	r.Post("/htmx/sign-in", handler.HTMX.SignInAction)
	r.Post("/htmx/sign-out", handler.HTMX.SignOut)

	r.Get("/htmx/sign-up", handler.HTMX.HomePage)
	r.Post("/htmx/sign-up", handler.HTMX.HomePage)

	r.Get("/htmx/home", handler.HTMX.HomePage)
	r.Get("/htmx/home/tag-list", handler.HTMX.HomeTagList)
	r.Get("/htmx/home/global-feed", handler.HTMX.HomeGlobalFeed)
	r.Get("/htmx/articles/:slug", handler.HTMX.HomeArticleDetailPage)
}

func registerHTMXCheckRouter(r fiber.Router) {
}
