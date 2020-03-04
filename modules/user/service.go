package user

import (
	"github.com/davidchandra95/go-laundry/model"
)

type Service interface {
	GetUsers() (interface{}, error)
	GetUser(id int64) (interface{}, error)
	AddUser(user model.User) (interface{}, error)
}
