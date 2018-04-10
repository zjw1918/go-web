package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string
	Email string	`gorm:"type:varchar(100);unique_index" json:"email"`
}

type UserModel struct {}

func (m UserModel) Signup() (user User, err error) {
	
}

func (m UserModel) Login() (user User, err error) {

}

func One()  {
	
}



