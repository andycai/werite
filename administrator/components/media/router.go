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

func initRootCheckRouter(r fiber.Router) {
	g := r.Group("/media/")
	g.Get("/*", handleMedia)
}

func initAdminCheckRouter(r fiber.Router) {
	r.Get("/media/manager", handleManagerPage)

	r.Post("/media/upload", handleUpload)
	r.Post("/media/delete", handleDelete)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterRootNoCheckRouter(KeyNoCheckRouter, initRootCheckRouter)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initAdminCheckRouter)
}
