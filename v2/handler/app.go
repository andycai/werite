package handler

import (
	"github.com/andycai/werite/v2/model"
	"github.com/gofiber/fiber/v2"
)

type AppHandler struct{}

var App = new(AppHandler)

//#region private methods

//

//#endregion

//#region handlers

// HomePage render home page
func (ah AppHandler) HomePage(c *Ctx) error {
	var authenticatedUser model.User

	// isAuthenticated, userID := authentication.AuthGet(c)

	return render(c, "home/index", fiber.Map{
		"PageTitle":         "Andy's Blog Home Page",
		"FiberCtx":          c,
		"NavBarActive":      "home",
		"AuthenticatedUser": authenticatedUser,
		"CurrentPage":       c.QueryInt("page"),
	}, "layouts/app")
}

// HomePage render home page
func (ah AppHandler) Article(c *Ctx) error {
	return nil
}

// HomePage render home page
func (ah AppHandler) Page(c *Ctx) error {
	return nil
}

// HomePage render home page
func (ah AppHandler) Series(c *Ctx) error {
	return nil
}

// HomePage render home page
func (ah AppHandler) Archive(c *Ctx) error {
	return nil
}

// HomePage render home page
func (ah AppHandler) Search(c *Ctx) error {
	return nil
}

//#endregion
