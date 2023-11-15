package handler

import (
	"errors"

	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/v2/model"
	"github.com/andycai/werite/v2/system"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SignInPage(c *fiber.Ctx) error {

	isAuthenticated, _ := authentication.AuthGet(c)
	if isAuthenticated {
		return c.Redirect("/")
	}

	return Render(c, "sign-in/index", fiber.Map{
		"PageTitle":    "Sign In — Werite",
		"FiberCtx":     c,
		"NavBarActive": "sign-in",
	}, "layouts/app")
}

//#region HMTX interface

func HTMXSignInPage(c *fiber.Ctx) error {
	return Render(c, "sign-in/htmx-sign-in-page", fiber.Map{
		"PageTitle":    "Sign In",
		"NavBarActive": "sign-in",
		"FiberCtx":     c,
	}, "layouts/app-htmx")
}

func HTMXSignInAction(c *fiber.Ctx) error {
	var user model.User
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return Render(c, "sign-in/partials/sign-in-form", fiber.Map{
			"Errors": []string{
				"Email or password cannot be null.",
			},
			"IsOob": true,
		}, "layouts/app-htmx")
	}

	err := system.User.FindByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Render(c, "sign-in/partials/sign-in-form", fiber.Map{
				"Errors": []string{
					"Email and password did not match.",
				},
			}, "layouts/app-htmx")
		}
	}

	if !CheckPassword(user.Password, password) {
		return Render(c, "sign-in/partials/sign-in-form", fiber.Map{
			"Errors": []string{
				"Email and password did not match.",
			},
		}, "layouts/app-htmx")
	}

	authentication.AuthStore(c, uint(user.ID))

	return HTMXRedirectTo("/", "/htmx/home", c)
}

func HTMXSignOut(c *fiber.Ctx) error {
	isAuthenticated, _ := authentication.AuthGet(c)
	if !isAuthenticated {
		return c.Redirect("/")
	}

	authentication.AuthDestroy(c)

	return HTMXRedirectTo("/", "/htmx/home", c)
}

//#endregion
