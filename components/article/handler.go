package article

import (
	"errors"

	"github.com/andycai/werite/components/user"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/v2/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ArticleDetailPage(c *fiber.Ctx) error {
	var article *model.Article
	var authenticatedUser model.User
	isSelf := false

	isAuthenticated, userID := authentication.AuthGet(c)

	article, err := Dao.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	if isAuthenticated {
		authenticatedUser = *user.Dao.GetByID(userID)
	}

	return core.Render(c, "articles/show", fiber.Map{
		"PageTitle":          article.Title + " — Werite",
		"Article":            article,
		"FiberCtx":           c,
		"IsOob":              false,
		"IsSelf":             isSelf,
		"IsArticleFavorited": false,
		"AuthenticatedUser":  authenticatedUser,
	}, "layouts/app")
}

//#region HTMX interface

// HTMXHomeArticleDetailPage detail page
func HTMXHomeArticleDetailPage(c *fiber.Ctx) error {
	var article *model.Article
	isSelf := false
	var authenticatedUser *model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		authenticatedUser = user.Dao.GetByID(userID)
	}

	article, err := Dao.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	return core.Render(c, "articles/htmx-article-page", fiber.Map{
		"PageTitle":          article.Title,
		"NavBarActive":       "none",
		"Article":            article,
		"IsOob":              false,
		"IsSelf":             isSelf,
		"IsArticleFavorited": false,
		"AuthenticatedUser":  authenticatedUser,
		"FiberCtx":           c,
	}, "layouts/app-htmx")
}

//#endregion
