package repository

import (
	"github.com/davidchandra95/go-laundry/model"
	"github.com/davidchandra95/go-laundry/modules/user"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) user.Repository {
	return &userRepository{Conn}
}

func (u *userRepository) GetUsers() ([]model.User, error) {
	var (
		err   error
		users []model.User
	)

	if err = u.Conn.Find(&users).Error; err != nil {
		logrus.Error(err)
		return users, err
	}

	return users, err
}

func (u *userRepository) GetUser(id int64) (model.User, error) {
	var (
		user model.User
		err  error
	)

	if err = u.Conn.First(&user, id).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, err
}

func (u *userRepository) Insert(newUser model.User) (model.User, error) {
	var (
		err  error
		user model.User
	)

	if err = u.Conn.Table("users").Create(&newUser).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, err
}
