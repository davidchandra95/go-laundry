package model

import "github.com/jinzhu/gorm"

type User struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      int    `json:"role"`
	gorm.Model
}
