package system

import "github.com/andycai/werite/v2/model"

type UserSystem struct {
}

var User = new(UserSystem)

func (us UserSystem) GetByID(id uint) *model.User {
	var user model.User
	db.Model(&user).
		Where("id = ?", id).
		First(&user)

	return &user
}

func (us UserSystem) FindByEmail(email string) error {
	var user model.User
	db.Model(&user)
	err := db.Where(&model.User{Email: email}).
		First(&user).Error

	return err
}

func (us UserSystem) Create(user *model.User) error {
	db.Create(user)

	return nil
}
