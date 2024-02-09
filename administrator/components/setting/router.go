package setting

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyDB            = "admin.setting.gorm.db"
	KeyNoCheckRouter = "admin.setting.router.nocheck"
	KeyCheckRouter   = "admin.setting.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initCheckRouter(r fiber.Router) {
	r.Get("/settings/blog", handleBlogPage)
	r.Post("/settings/blog", handleBlogSave)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
