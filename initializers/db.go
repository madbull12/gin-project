package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	// DB, err := gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic("Failed to connect to ddb")
	}

}
