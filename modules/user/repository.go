package user

import "github.com/davidchandra95/go-laundry/model"

type Repository interface {
	GetUsers() ([]model.User, error)
	GetUser(id int64) (model.User, error)
	Insert(user model.User) (model.User, error)
}
