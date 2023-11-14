package handler

import (
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/library/database"
	"github.com/andycai/werite/v2/model"
	"github.com/gofiber/fiber/v2"
)

func SignUpPage(c *fiber.Ctx) error {

	isAuthenticated, _ := authentication.AuthGet(c)
	if isAuthenticated {
		return c.Redirect("/")
	}

	return render(c, "sign-up/index", fiber.Map{
		"PageTitle":    "Sign Up — Werite",
		"FiberCtx":     c,
		"NavBarActive": "sign-up",
	}, "layouts/app")
}

func HTMXSignUpPage(c *fiber.Ctx) error {
	return c.Render("sign-up/htmx-sign-up-page", fiber.Map{
		"PageTitle":    "Sign Up",
		"NavBarActive": "sign-up",
		"FiberCtx":     c,
	}, "layouts/app-htmx")
}

func HTMXSignUpAction(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || username == "" || password == "" {

		return c.Render("sign-up/partials/sign-up-form", fiber.Map{
			"Errors": []string{
				"Username, email, and password cannot be null.",
			},
			"IsOob": true,
		}, "layouts/app-htmx")
	}

	user := model.User{Username: username, Email: email, Password: password, Name: username}
	user.Password = HashPassword(user.Password)

	db := database.Get()
	db.Create(&user)

	authentication.AuthStore(c, uint(user.ID))

	return HTMXRedirectTo("/", "/htmx/home", c)
}
