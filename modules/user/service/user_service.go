package service

import (
	"github.com/davidchandra95/go-laundry/model"
	"github.com/davidchandra95/go-laundry/modules/user"
	"github.com/sirupsen/logrus"
)

type userService struct {
	userRepo    user.Repository
}

func NewUserService(u user.Repository) user.Service {
	return &userService{
		userRepo: u,
	}
}

func (u *userService) GetUsers() (interface{}, error) {
	var (
		err error
		users []model.User
	)

	if users, err = u.userRepo.GetUsers(); err != nil {
		logrus.Error(err)
		return users, err
	}

	return users, err
}

func (u *userService) GetUser(id int64) (interface{}, error) {
	var (
		err error
		res model.User
	)

	if res, err = u.userRepo.GetUser(id); err != nil {
		logrus.Error(err)
		return res, err
	}

	return res, err
}

func (u *userService) AddUser(newUser model.User) (interface{}, error) {
	var (
		err error
		res model.User
	)

	if res, err = u.userRepo.Insert(newUser); err != nil {
		logrus.Error(err)
		return res, err
	}

	return res, err
}