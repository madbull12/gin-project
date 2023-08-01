package config

import (
	"github.com/madbull12/gin_postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("user=postgres password=Qp83@q!9Cz!Ghni host=db.eneqclqeabkawrxopsxn.supabase.co port=5432 dbname=postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db

}
