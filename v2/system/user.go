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
