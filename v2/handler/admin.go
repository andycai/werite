package handler

import "github.com/gofiber/fiber/v2"

type AdminHandler struct{}

var Admin = new(AdminHandler)

//#region private methods

//

//#endregion

//#region handler

// Login login to admin
func (ah AdminHandler) Login(c *Ctx) error {
	return Render(c, "admin/login", fiber.Map{})
}

//#endregion
