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
	r.Get("/posts/manager", handleManagerPage)
	r.Get("/posts/editor", handleEditorPage)
	r.Get("/posts/editor/:id", handleEditorPage)
	r.Post("/posts/editor", handleCreate)
	r.Post("/posts/editor/:id", handleUpdate)
	r.Get("/posts/movetotrash/:id", handleMoveToTrashByID)
	r.Post("/posts/movetotrash", handleMoveToTrash)
	r.Get("/posts/restore/:id", handleRestoreByID)
	r.Post("/posts/restore", handleRestore)

	r.Get("/categories/manager", handleManagerCategoryPage)
	r.Get("/categories/editor", handleEditorCategoryPage)
	r.Get("/categories/editor/:id", handleEditorCategoryPage)
	r.Post("/categories/editor", handleCreateCategory)
	r.Post("/categories/editor/:id", handleUpdateCategory)
	r.Post("/categories/delete", handleDeleteCategories)

	r.Get("/tags/manager", handleManagerTagsPage)
	r.Get("/tags/editor", handleEditorTagPage)
	r.Post("/tags/editor", handleCreateTag)
	r.Post("/tags/delete", handleDeleteTags)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
