package user

import "github.com/andycai/werite/components/user/model"

type UserDao struct{}

var Dao = new(UserDao)

func (ud UserDao) GetByID(id uint) *model.User {
	var user model.User
	db.Model(&user).
		Where("id = ?", id).
		First(&user)

	return &user
}

func (ud UserDao) FindByEmail(email string) (error, model.User) {
	var user model.User
	db.Model(&user)
	err := db.Where(&model.User{Email: email}).
		First(&user).Error

	return err, user
}

func (ud UserDao) Create(user *model.User) error {
	db.Create(user)

	return nil
}
