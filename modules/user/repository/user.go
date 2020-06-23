package repository

import (
	"errors"
	"github.com/davidchandra95/go-laundry/model"
	"github.com/davidchandra95/go-laundry/modules/user"
	"github.com/davidchandra95/go-laundry/util"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) user.Repository {
	return &userRepository{Conn}
}

func (r *userRepository) GetUsers() ([]model.User, error) {
	var (
		err   error
		users []model.User
	)

	if err = r.Conn.Find(&users).Error; err != nil {
		logrus.Error(err)
		return users, err
	}

	return users, err
}

func (r *userRepository) GetUser(id int64) (model.User, error) {
	var (
		user model.User
		err  error
	)

	if err = r.Conn.First(&user, id).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, err
}

func (r *userRepository) Insert(newUser model.User) (model.User, error) {
	var (
		err  error
		user model.User
	)

	if err = r.Conn.Table("users").Create(&newUser).Error; err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, err
}

func (r *userRepository) Update(uid uint, newUser model.User) (*model.User, error) {
	var existingUser *model.User

	if err := r.Conn.Where("id = ?", uid).Find(&existingUser).Error; err != nil {
		logrus.Error(err)
		return existingUser, err
	}

	hashedPass, err := util.HashPassword(newUser.Password)
	if err != nil {
		logrus.Error(err)
		return existingUser, err
	}

	existingUser.Email = newUser.Email
	existingUser.Username = newUser.Username
	existingUser.Password = hashedPass
	existingUser.UpdatedAt = time.Now()

	if err := r.Conn.Save(&existingUser).Error; err != nil {
		logrus.Error(err)
		return existingUser, err
	}

	return existingUser, err
}

func (r *userRepository) Delete(uid uint) (int64, error) {
	res := r.Conn.Where("id = ?", uid).Take(&model.User{}).Delete(&model.User{})
	if res.Error != nil {
		logrus.Error(res.Error)
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (r *userRepository) Validate(action string, user model.User) error {
	switch strings.ToLower(action) {
	case "update":
		if user.Username == "" {
			return errors.New("username required")
		}
		if user.Password == "" {
			return errors.New("password required")
		}
		if user.Email == "" {
			return errors.New("email required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}

		return nil
	case "login":
		if user.Password == "" {
			return errors.New("password required")
		}
		if user.Email == "" {
			return errors.New("email required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if user.Username == "" {
			return errors.New("username required")
		}
		if user.Password == "" {
			return errors.New("password required")
		}
		if user.Email == "" {
			return errors.New("email required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email")
		}

		return nil
	}
}
