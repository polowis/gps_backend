package models

import (
	"github.com/gps/conf"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model      // standard fields
	Email 	 string // user email
	Password string // box order must be hashed
	Casting  string // coordinate drawing must be encoded
}

func (u *User) Create() error {
	result := conf.DB.Create(&u)
	return result.Error
}