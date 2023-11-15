package handler

import (
	"errors"

	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/v2/model"
	"github.com/andycai/werite/v2/system"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ArticleDetailPage(c *Ctx) error {
	var article *model.Article
	var authenticatedUser model.User
	isSelf := false

	isAuthenticated, userID := authentication.AuthGet(c)

	article, err := system.Article.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	if isAuthenticated {
		authenticatedUser = *system.User.GetByID(userID)
	}

	return Render(c, "articles/show", fiber.Map{
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
func HTMXHomeArticleDetailPage(c *Ctx) error {
	var article *model.Article
	isSelf := false
	var authenticatedUser *model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		authenticatedUser = system.User.GetByID(userID)
	}

	article, err := system.Article.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	return Render(c, "articles/htmx-article-page", fiber.Map{
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
