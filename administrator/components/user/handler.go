package user

import (
	"errors"
	"fmt"

	"github.com/andycai/werite/components/page"
	"github.com/andycai/werite/components/post"
	"github.com/andycai/werite/components/user"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginPage(c *fiber.Ctx) error {
	isAuthenticated, _ := authentication.AuthGet(c)

	if isAuthenticated {
		return c.Redirect("/admin/dashboard")
	}

	return core.Render(c, "admin/login", fiber.Map{})
}

func LoginAction(c *fiber.Ctx) error {
	var userVo *model.User
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

	user.Dao.UpdateLoginTime(uint(userVo.ID))
	authentication.AuthStore(c, uint(userVo.ID))

	return c.Redirect("/admin/dashboard")
}

func handleLogoutAction(c *fiber.Ctx) error {
	isAuthenticated, userID := authentication.AuthGet(c)
	if !isAuthenticated {
		return core.Render(c, "admin/login", fiber.Map{})
	}

	user.Dao.UpdateLogoutTime(userID)
	authentication.AuthDestroy(c)

	return core.Render(c, "admin/login", fiber.Map{})
}

func handleDashBoardPage(c *fiber.Ctx) error {
	userTotal := user.Dao.Count()
	postTotal := post.Dao.Count()
	pageTotal := page.Dao.Count()

	bind := fiber.Map{
		"PageTitle":    "DashBoard",
		"NavBarActive": "dashboard",
		"Path":         "/admin/dashboard",
		"UserTotal":    userTotal,
		"PostTotal":    postTotal,
		"PageTotal":    pageTotal,
	}
	DecorateUserBar(c, bind)

	return core.Render(c, "admin/dashboard", bind, "admin/layouts/app")
}

func handleProfilePage(c *fiber.Ctx) error {
	var userVo *model.User
	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		userVo = user.Dao.GetByID(userID)
	}

	bind := fiber.Map{
		"PageTitle":    "Profile",
		"NavBarActive": "users",
		"Path":         "/admin/users/profile",
		"User":         userVo,
	}
	DecorateUserBar(c, bind)

	return core.Render(c, "admin/users/profile", bind, "admin/layouts/app")
}

func handleSecurityPage(c *fiber.Ctx) error {
	bind := fiber.Map{
		"PageTitle":    "Profile",
		"NavBarActive": "users",
		"Path":         "/admin/users/security",
	}
	DecorateUserBar(c, bind)

	return core.Render(c, "admin/users/security", bind, "admin/layouts/app")
}

func handleProfileSave(c *fiber.Ctx) error {
	var userVo *model.User
	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		userVo = user.Dao.GetByID(userID)
	}

	err := user.BindProfile(c, userVo)
	if err != nil {
		return err
	}

	db.Model(userVo).Updates(map[string]interface{}{"gender": userVo.Gender, "phone": userVo.Phone, "email": userVo.Email, "addr": userVo.Addr})

	core.PushMessages(fmt.Sprintf("Updated profile"))

	return c.Redirect("/admin/users/profile")
}

func handlePasswordSave(c *fiber.Ctx) error {
	var userVo *model.User
	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		userVo = user.Dao.GetByID(userID)
	}

	err := user.BindPassword(c, userVo)
	if err != nil {
		return err
	}

	db.Model(userVo).Update("password", userVo.Password)

	core.PushMessages(fmt.Sprintf("Updated password"))

	return c.Redirect("/admin/users/security")
}
