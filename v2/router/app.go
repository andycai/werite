package router

import (
	"github.com/andycai/werite/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerCheck = append(routerCheck, registerAppCheckRouter)
	routerNoCheck = append(routerNoCheck, registerAppNoCheckRouter)
}

func registerAppNoCheckRouter(r fiber.Router) {
	r.Get("/sign-in", handler.SignInPage)
	r.Get("/sign-up", handler.SignUpPage)

	r.Get("/", handler.HomePage)
	r.Get("/articles/:slug", handler.ArticleDetailPage)
	// r.Get("/page/:slug", handler.App.Page)
	// r.Get("/series.html", handler.App.Series)
	// r.Get("/archives.html", handler.App.Archive)
	// r.Get("/search.html", handler.App.Search)
}

func registerAppCheckRouter(r fiber.Router) {
}
