package setting

import (
	"github.com/andycai/werite/components/user"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
)

func BlogPage(c *fiber.Ctx) error {
	var userVo *model.User
	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		userVo = user.Dao.GetByID(userID)
	}

	return core.Render(c, "admin/settings/blog", fiber.Map{
		"PageTitle":    "Blog",
		"NavBarActive": "settings",
		"Path":         "/admin/settings/blog",
		"UserName":     userVo.Name,
		"Info": fiber.Map{
			"BlogName":     "Werite",
			"BlogSubTitle": "Content Management System",
			"LoginAt":      userVo.LoginAt,
		},
	}, "admin/layouts/app")
}

func BlogSave(c *fiber.Ctx) error {
	var blogVo *model.Blog

	err := user.BindBlog(c, blogVo)
	if err != nil {
		return err
	}

	db.Model(blogVo).Updates(map[string]interface{}{"name": blogVo.Name, "description": blogVo.Description})

	return c.Redirect("/admin/settings/blog")
}
