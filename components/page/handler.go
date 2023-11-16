package page

import (
	"errors"

	"github.com/andycai/werite/components/page/model"
	"github.com/andycai/werite/components/user"
	mu "github.com/andycai/werite/components/user/model"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PageDetailPage(c *fiber.Ctx) error {
	var page *model.Page
	var authenticatedUser mu.User

	isAuthenticated, userID := authentication.AuthGet(c)

	page, err := Dao.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	if isAuthenticated {
		authenticatedUser = *user.Dao.GetByID(userID)
	}

	return core.Render(c, "pages/show", fiber.Map{
		"PageTitle":         page.Title + " — Werite",
		"Page":              page,
		"FiberCtx":          c,
		"AuthenticatedUser": authenticatedUser,
	}, "layouts/app")
}

//#region HTMX interface

func HTMXHomePageDetailPage(c *fiber.Ctx) error {
	var page *model.Page
	var authenticatedUser *mu.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		authenticatedUser = user.Dao.GetByID(userID)
	}

	page, err := Dao.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	return core.Render(c, "pages/htmx--page", fiber.Map{
		"PageTitle":         page.Title,
		"NavBarActive":      "none",
		"Page":              page,
		"AuthenticatedUser": authenticatedUser,
		"FiberCtx":          c,
	}, "layouts/app-htmx")
}

//#endregion
