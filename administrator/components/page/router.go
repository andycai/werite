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

func initCheckRouter(r fiber.Router) {
	r.Get("/pages/manager", handleManagerPage)

	r.Get("/pages/editor", handleEditorPage)
	r.Get("/pages/editor/:id", handleEditorPage)
	r.Post("/pages/editor", handleCreate)
	r.Post("/pages/editor/:id", handleUpdate)
	r.Get("/pages/movetotrash/:id", handleMoveToTrashByID)
	r.Post("/pages/movetotrash", handleMoveToTrash)
	r.Get("/pages/restore/:id", handleRestoreByID)
	r.Post("/pages/restore", handleRestore)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
