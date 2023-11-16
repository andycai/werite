package user

import (
	"time"

	"github.com/andycai/werite/components/user/model"
)

type UserDao struct{}

var Dao = new(UserDao)

func (ud UserDao) GetByID(id uint) *model.User {
	var user model.User
	db.Model(&user).
		Where("id = ?", id).
		First(&user)

	return &user
}

func (ud UserDao) FindByEmail(email string) (error, *model.User) {
	var user model.User
	db.Model(&user)
	err := db.Where(&model.User{Email: email}).
		First(&user).Error

	return err, &user
}

func (ud UserDao) Create(user *model.User) error {
	db.Create(user)

	return nil
}

func (ud UserDao) UpdateLoginTime(userID uint) error {
	db.Model(&model.User{}).Where("id = ?", userID).Update("login_at", time.Now())

	return nil
}

func (ud UserDao) UpdateLogoutTime(userID uint) error {
	db.Model(&model.User{}).Where("id = ?", userID).Update("logout_at", time.Now())

	return nil
}
