package setting

import (
	"github.com/andycai/werite/components/user"
	"github.com/andycai/werite/components/user/model"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/gofiber/fiber/v2"
)

func ProfilePage(c *fiber.Ctx) error {
	var authenticatedUser *model.User
	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		authenticatedUser = user.Dao.GetByID(userID)
	}

	return core.Render(c, "admin/settings/profile", fiber.Map{
		"PageTitle": "Profile",
		"Path":      "admin/profile",
		"UserName":  authenticatedUser.Name,
		"Info": fiber.Map{
			"BlogName":     "Werite",
			"BlogSubTitle": "Content Management System",
			"LoginAt":      authenticatedUser.LoginAt,
		},
	}, "admin/layouts/app")
}

func ProfileAction(c *fiber.Ctx) error {
	return nil
}
