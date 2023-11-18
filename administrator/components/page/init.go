package page

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyDB            = "admin.page.gorm.db"
	KeyNoCheckRouter = "admin.page.router.nocheck"
	KeyCheckRouter   = "admin.page.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	// admin := r.Group("/admin")
	// {
	// }
}

func initCheckRouter(r fiber.Router) {
	admin := r.Group("/admin")
	{
		admin.Get("/pages", PagesPage)
		admin.Get("/page", PagePage)

		admin.Get("/page/:slug", EditorPage)
		admin.Post("/page", PageAction)
	}
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterNoCheckRouter(KeyNoCheckRouter, initNoCheckRouter)
	core.RegisterCheckRouter(KeyCheckRouter, initCheckRouter)
}
