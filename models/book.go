package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
