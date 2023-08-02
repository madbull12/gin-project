package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"password" validate:"required"`
}
