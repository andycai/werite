package handler

import (
	"errors"

	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/v2/model"
	"github.com/andycai/werite/v2/system"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PageDetailPage(c *Ctx) error {
	var page *model.Page
	var authenticatedUser model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	page, err := system.Page.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	if isAuthenticated {
		authenticatedUser = *system.User.GetByID(userID)
	}

	return Render(c, "pages/show", fiber.Map{
		"PageTitle":         page.Title + " — Werite",
		"Page":              page,
		"FiberCtx":          c,
		"AuthenticatedUser": authenticatedUser,
	}, "layouts/app")
}

//#region HTMX interface

func HTMXHomePageDetailPage(c *Ctx) error {
	var page *model.Page
	var authenticatedUser *model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		authenticatedUser = system.User.GetByID(userID)
	}

	page, err := system.Page.GetBySlug(c.Params("slug"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	return Render(c, "pages/htmx--page", fiber.Map{
		"PageTitle":         page.Title,
		"NavBarActive":      "none",
		"Page":              page,
		"AuthenticatedUser": authenticatedUser,
		"FiberCtx":          c,
	}, "layouts/app-htmx")
}

//#endregion
