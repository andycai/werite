package handler

import (
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/v2/model"
	"github.com/andycai/werite/v2/system"
	"github.com/gofiber/fiber/v2"
)

func SignUpPage(c *fiber.Ctx) error {
	isAuthenticated, _ := authentication.AuthGet(c)
	if isAuthenticated {
		return c.Redirect("/")
	}

	return Render(c, "sign-up/index", fiber.Map{
		"PageTitle":    "Sign Up — Werite",
		"FiberCtx":     c,
		"NavBarActive": "sign-up",
	}, "layouts/app")
}

//#region HTMX interface

func HTMXSignUpPage(c *fiber.Ctx) error {
	return Render(c, "sign-up/htmx-sign-up-page", fiber.Map{
		"PageTitle":    "Sign Up",
		"NavBarActive": "sign-up",
		"FiberCtx":     c,
	}, "layouts/app-htmx")
}

func HTMXSignUpAction(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// TODO: validate
	if email == "" || username == "" || password == "" {
		return Render(c, "sign-up/partials/sign-up-form", fiber.Map{
			"Errors": []string{
				"Username, email, and password cannot be null.",
			},
			"IsOob": true,
		}, "layouts/app-htmx")
	}

	user := model.User{Username: username, Email: email, Password: password, Name: username}
	user.Password = HashPassword(user.Password)

	err := system.User.Create(&user)

	if err != nil {
		return err
	}

	authentication.AuthStore(c, uint(user.ID))

	return HTMXRedirectTo("/", "/htmx/home", c)
}

//#endregion
