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

func initNoCheckRouter(r fiber.Router) {
}

func initCheckRouter(r fiber.Router) {
	admin := r.Group("/admin")
	{
		admin.Get("/profile", ProfilePage)
		admin.Post("/profile", ProfileAction)
	}
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterNoCheckRouter(KeyNoCheckRouter, initNoCheckRouter)
	core.RegisterCheckRouter(KeyCheckRouter, initCheckRouter)
}
