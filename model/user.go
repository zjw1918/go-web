package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zjw1918/go-web/forms"
	"github.com/zjw1918/go-web/db"
	"errors"
	"github.com/zjw1918/go-web/utils"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string	`json:"-"`
	Email string	`gorm:"type:varchar(100);unique_index" json:"email"`
}

//UserSessionInfo ...
type UserSessionInfo struct {
	ID    int64  `json:"id"`
	Username  string `json:"username"`
	Email string `json:"email"`
}

type UserModel struct {}

func (m UserModel) Signup(form forms.SignupForm) (user User, err error) {
	var users []User
	err = db.GetDB().First(&users, "username = ?", form.Username).Error
	log.Println(users)
	if err != nil {
		return user, nil
	}

	if len(users) > 0 {
		return user, errors.New("user or email already exists")
	}

	err = db.GetDB().First(&users, "email = ?", form.Email).Error
	log.Println(users)
	if err != nil {
		return user, nil
	}

	if len(users) > 0 {
		return user, errors.New("user or email already exists")
	}

	hash, _ := utils.HashPassword(form.Password)
	res := User{
		Username: form.Username,
		Password: hash,
		Email: form.Email,
	}
	err = db.GetDB().Create(&res).Error
	if err != nil {
		return user, err
	}
	return res, nil
}

func (m UserModel) Signin(form forms.SigninForm) (user User, err error) {
	var users []User
	err = db.GetDB().First(&users, "username = ?", form.Username).Error

	log.Println(users)
	if err != nil {
		return user, err
	}

	if len(users) <= 0 || !utils.CheckHashPassword(users[0].Password, form.Password) {
		return user, errors.New("username or password wrong")
	}
	return users[0], err
}

func (m UserModel) All() (users []User, err error) {
	err = db.GetDB().Find(&users).Error
	return users, err
}



