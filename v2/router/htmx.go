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
	r.Get("/htmx/sign-in", handler.HTMXSignInPage)
	r.Post("/htmx/sign-in", handler.HTMXSignInAction)
	r.Post("/htmx/sign-out", handler.HTMXSignOut)

	r.Get("/htmx/sign-up", handler.HTMXSignUpPage)
	r.Post("/htmx/sign-up", handler.HTMXSignUpAction)

	r.Get("/htmx/home", handler.HTMXHomePage)
	r.Get("/htmx/home/tag-list", handler.HTMXHomeTagList)
	r.Get("/htmx/home/global-feed", handler.HTMXHomeGlobalFeed)
	r.Get("/htmx/articles/:slug", handler.HTMXHomeArticleDetailPage)
}

func registerHTMXCheckRouter(r fiber.Router) {
}
