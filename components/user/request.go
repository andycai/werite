package user

import (
	"errors"

	"github.com/andycai/werite/core"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
)

type requestChangingPassword struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=NewPassword"`
}

func BindPassword(c *fiber.Ctx, user *model.User) error {
	var r requestChangingPassword
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	if err := core.Validate(r); err != nil {
		return err
	}

	if !core.CheckPassword(user.Password, r.CurrentPassword) {
		return errors.New("Current password is error")
	}

	user.Password = core.HashPassword(r.NewPassword)

	return nil
}
