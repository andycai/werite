package post

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyPostDB            = "post.gorm.db"
	KeyPostNoCheckRouter = "post.router.nocheck"
	KeyPostCheckRouter   = "post.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	r.Get("/posts/:slug", DetailPage)

	// HTMX
	r.Get("/htmx/posts/:slug", HTMXHomePostDetailPage)
}

func initCheckRouter(r fiber.Router) {
	//
}

func init() {
	core.RegisterDatabase(KeyPostDB, initDB)
	core.RegisterRootNoCheckRouter(KeyPostNoCheckRouter, initNoCheckRouter)
	core.RegisterAPICheckRouter(KeyPostCheckRouter, initCheckRouter)
}
