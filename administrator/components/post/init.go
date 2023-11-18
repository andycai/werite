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

func initNoCheckRouter(r fiber.Router) {
}

func initCheckRouter(r fiber.Router) {
	admin := r.Group("/admin")
	{
		admin.Get("/posts/manager", ManagerPage)

		admin.Get("/posts/editor", EditorPage)
		admin.Get("/posts/editor/:slug", EditorPage)
		admin.Post("/posts/editor", EditorAction)
	}
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterNoCheckRouter(KeyNoCheckRouter, initNoCheckRouter)
	core.RegisterCheckRouter(KeyCheckRouter, initCheckRouter)
}
