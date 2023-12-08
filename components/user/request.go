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

type requestProfile struct {
	Gender uint   `json:"gender" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
	Email  string `json:"email" validate:"email"`
	Addr   string `json:"addr" validate:"required"`
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

func BindProfile(c *fiber.Ctx, user *model.User) error {
	var r requestProfile
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	if err := core.Validate(r); err != nil {
		return err
	}

	user.Gender = r.Gender
	user.Phone = r.Phone
	user.Email = r.Email
	user.Addr = r.Addr

	return nil
}