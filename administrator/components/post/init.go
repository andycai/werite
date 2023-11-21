package post

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyDB            = "admin.post.gorm.db"
	KeyNoCheckRouter = "admin.post.router.nocheck"
	KeyCheckRouter   = "admin.post.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initCheckRouter(r fiber.Router) {
	r.Get("/posts/manager", ManagerPage)

	r.Get("/posts/editor", EditorPage)
	r.Get("/posts/editor/:id", EditorPage)
	r.Post("/posts/editor", Create)
	r.Post("/posts/editor/:id", Update)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
