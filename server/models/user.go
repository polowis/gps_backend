package models

import (
	"github.com/gps/conf"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email 	 string
	Password string
	Casting  string
}

func (u *User) Create() error {
	result := conf.DB.Create(&u)
	return result.Error
}