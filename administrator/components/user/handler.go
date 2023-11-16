package user

import (
	"errors"
	"time"

	"github.com/andycai/werite/components/user"
	"github.com/andycai/werite/components/user/model"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginPage(c *fiber.Ctx) error {
	isAuthenticated, _ := authentication.AuthGet(c)

	if isAuthenticated {
		return c.Redirect("/admin/profile")
	}

	return core.Render(c, "admin/login", fiber.Map{})
}

func LoginAction(c *fiber.Ctx) error {
	var userVo model.User
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return core.Render(c, "admin/login", fiber.Map{
			"Errors": []string{
				"Email or password cannot be null.",
			},
		})
	}

	err, userVo := user.Dao.FindByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return core.Render(c, "admin/login", fiber.Map{
				"Errors": []string{
					"Email and password did not match.",
				},
			})
		}
	}

	if !core.CheckPassword(userVo.Password, password) {
		return core.Render(c, "admin/login", fiber.Map{
			"Errors": []string{
				"Email and password did not match.",
			},
		})
	}

	authentication.AuthStore(c, uint(userVo.ID))

	return core.Render(c, "admin/profile", fiber.Map{})
}

func ProfilePage(c *fiber.Ctx) error {
	// isAuthenticated, _ := authentication.AuthGet(c)

	// if isAuthenticated {
	// 	return core.Render(c, "admin/profile", fiber.Map{})
	// }

	return core.Render(c, "admin/profile", fiber.Map{
		"PageTitle": "Profile",
		"Path":      "admin/profile",
		"Console":   true,
		"Profile": fiber.Map{
			"BlogName":     "Werite",
			"BlogSubTitle": "Content Management System",
			"LoginAt":      time.Now(),
		},
	}, "admin/layouts/app")
}

func ProfileAction(c *fiber.Ctx) error {
	return nil
}
