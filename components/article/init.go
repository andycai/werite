package article

import (
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyArticleDB            = "article.gorm.db"
	KeyArticleNoCheckRouter = "article.router.nocheck"
	KeyArticleCheckRouter   = "article.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	r.Get("/articles/:slug", ArticleDetailPage)

	// HTMX
	r.Get("/htmx/articles/:slug", HTMXHomeArticleDetailPage)
}

func initCheckRouter(r fiber.Router) {
	//
}

func init() {
	core.RegisterDatabase(KeyArticleDB, initDB)
	core.RegisterNoCheckRouter(KeyArticleNoCheckRouter, initNoCheckRouter)
	core.RegisterCheckRouter(KeyArticleCheckRouter, initCheckRouter)
}
