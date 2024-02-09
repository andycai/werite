package media

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyDB            = "admin.media.gorm.db"
	KeyNoCheckRouter = "admin.media.router.nocheck"
	KeyCheckRouter   = "admin.media.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initCheckRouter(r fiber.Router) {
	r.Get("/media/manager", handleManagerPage)

	r.Post("/media/upload", handleUpload)
	r.Post("/media/delete", handleDelete)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
