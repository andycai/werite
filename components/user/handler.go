package user

import (
	"errors"

	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/v2/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//#region sign in

func SignInPage(c *fiber.Ctx) error {

	isAuthenticated, _ := authentication.AuthGet(c)
	if isAuthenticated {
		return c.Redirect("/")
	}

	return core.Render(c, "sign-in/index", fiber.Map{
		"PageTitle":    "Sign In — Werite",
		"FiberCtx":     c,
		"NavBarActive": "sign-in",
	}, "layouts/app")
}

func HTMXSignInPage(c *fiber.Ctx) error {
	return core.Render(c, "sign-in/htmx-sign-in-page", fiber.Map{
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
		return core.Render(c, "sign-in/partials/sign-in-form", fiber.Map{
			"Errors": []string{
				"Email or password cannot be null.",
			},
			"IsOob": true,
		}, "layouts/app-htmx")
	}

	err := Dao.FindByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return core.Render(c, "sign-in/partials/sign-in-form", fiber.Map{
				"Errors": []string{
					"Email and password did not match.",
				},
			}, "layouts/app-htmx")
		}
	}

	if !core.CheckPassword(user.Password, password) {
		return core.Render(c, "sign-in/partials/sign-in-form", fiber.Map{
			"Errors": []string{
				"Email and password did not match.",
			},
		}, "layouts/app-htmx")
	}

	authentication.AuthStore(c, uint(user.ID))

	return core.HTMXRedirectTo("/", "/htmx/home", c)
}

func HTMXSignOut(c *fiber.Ctx) error {
	isAuthenticated, _ := authentication.AuthGet(c)
	if !isAuthenticated {
		return c.Redirect("/")
	}

	authentication.AuthDestroy(c)

	return core.HTMXRedirectTo("/", "/htmx/home", c)
}

//#endregion

//#region sign up

func SignUpPage(c *fiber.Ctx) error {
	isAuthenticated, _ := authentication.AuthGet(c)
	if isAuthenticated {
		return c.Redirect("/")
	}

	return core.Render(c, "sign-up/index", fiber.Map{
		"PageTitle":    "Sign Up — Werite",
		"FiberCtx":     c,
		"NavBarActive": "sign-up",
	}, "layouts/app")
}

func HTMXSignUpPage(c *fiber.Ctx) error {
	return core.Render(c, "sign-up/htmx-sign-up-page", fiber.Map{
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
		return core.Render(c, "sign-up/partials/sign-up-form", fiber.Map{
			"Errors": []string{
				"Username, email, and password cannot be null.",
			},
			"IsOob": true,
		}, "layouts/app-htmx")
	}

	user := model.User{Username: username, Email: email, Password: password, Name: username}
	user.Password = core.HashPassword(user.Password)

	err := Dao.Create(&user)

	if err != nil {
		return err
	}

	authentication.AuthStore(c, uint(user.ID))

	return core.HTMXRedirectTo("/", "/htmx/home", c)
}

//#endregion
