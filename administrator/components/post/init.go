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

	r.Get("/categories/manager", ManagerCategoryPage)
	r.Get("/categories/editor", EditorCategoryPage)
	r.Get("/categories/editor/:id", EditorCategoryPage)
	r.Post("/categories/editor", CreateCategory)
	r.Post("/categories/editor/:id", UpdateCategory)

	r.Get("/tags/manager", ManagerTagsPage)
	r.Get("/tags/editor", EditorTagPage)
	r.Get("/tags/editor/:id", EditorTagPage)
	r.Post("/tags/editor", CreateTag)
	r.Post("/tags/editor/:id", UpdateTag)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
