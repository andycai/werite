package setting

import (
	"errors"
	"fmt"

	adminuser "github.com/andycai/werite/administrator/components/user"
	"github.com/andycai/werite/components/user"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BlogPage(c *fiber.Ctx) error {
	_, userID := authentication.AuthGet(c)

	var blogVo model.Blog
	err := db.Model(blogVo).Where("user_id= ?", userID).First(&blogVo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		blogVo = model.Blog{}
	}

	bind := fiber.Map{
		"PageTitle":    "Blog",
		"NavBarActive": "settings",
		"Path":         "/admin/settings/blog",
		"Blog":         blogVo,
	}
	adminuser.DecorateUserBar(c, bind)

	return core.Render(c, "admin/settings/blog", bind, "admin/layouts/app")
}

func BlogSave(c *fiber.Ctx) error {
	blogVo := model.Blog{}

	_, userID := authentication.AuthGet(c)

	err := user.BindBlog(c, &blogVo)
	if err != nil {
		return err
	}

	blogVo.UserID = userID

	err = db.Model(&blogVo).Where("user_id= ?", userID).First(&blogVo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&blogVo)
	} else {
		db.Model(&blogVo).Updates(map[string]interface{}{"name": blogVo.Name, "description": blogVo.Description})
	}

	core.PushMessages(fmt.Sprintf("Updated blog infomation"))

	return c.Redirect("/admin/settings/blog")
}
