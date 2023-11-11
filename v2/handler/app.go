package handler

import (
	"github.com/gofiber/fiber/v2"
)

type AppHandler struct{}

var App = new(AdminHandler)

//#region private methods

//

//#endregion

//#region handlers

// HomePage render home page
func (uh AdminHandler) HomePage(c *Ctx) error {
	return Render(c, "index", fiber.Map{
		"Title": "Andy's Blog Home Page",
	})
}

// HomePage render home page
func (uh AdminHandler) Article(c *Ctx) error {
	return nil
}

// HomePage render home page
func (uh AdminHandler) Page(c *Ctx) error {
	return nil
}

// HomePage render home page
func (uh AdminHandler) Series(c *Ctx) error {
	return nil
}

// HomePage render home page
func (uh AdminHandler) Archive(c *Ctx) error {
	return nil
}

// HomePage render home page
func (uh AdminHandler) Search(c *Ctx) error {
	return nil
}

//#endregion
