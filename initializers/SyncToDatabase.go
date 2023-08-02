package initializers

import "github.com/madbull12/gin_postgres/models"

func SyncToDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
