package page

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyPageDB            = "page.gorm.db"
	KeyPageNoCheckRouter = "page.router.nocheck"
	KeyPageCheckRouter   = "page.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	r.Get("/pages/:slug", DetailPage)

	// HTMX

}

func initCheckRouter(r fiber.Router) {
	//
}

func init() {
	core.RegisterDatabase(KeyPageDB, initDB)
	core.RegisterNoCheckRouter(KeyPageNoCheckRouter, initNoCheckRouter)
	core.RegisterCheckRouter(KeyPageNoCheckRouter, initCheckRouter)
}
