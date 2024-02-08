package user

import (
	"github.com/andycai/werite/components/user"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
)

// DecorateUserBar decorates the user bar.
func DecorateUserBar(c *fiber.Ctx, bind fiber.Map) {
	var userVo *model.User
	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		userVo = user.Dao.GetByID(userID)
		bind["UserName"] = userVo.Name
		// bind["User"] = userVo
		bind["Info"] = fiber.Map{
			"AdminName":        "Werite",
			"AdminDescription": "Content Management System",
			"LoginAt":          userVo.LoginAt,
		}
	}
}
