package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/madbull12/gin_postgres/initializers"
	"github.com/madbull12/gin_postgres/models"
)

func GetBooks(c *gin.Context) {
	books := []models.Book{}
	initializers.DB.Find(&books)
	c.JSON(200, &books)
}

// func GetSingleUser(c *gin.Context) {
// 	user := models.User{}
// 	initializers.DB.Where("id = ?", c.Param("id")).First(&user)

// 	c.BindJSON(&user)
// 	initializers.DB.Find(&user)

// 	c.JSON(200, user)
// }

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	initializers.DB.Create(&book)
	c.JSON(200, &book)

}

func DeleteBook(c *gin.Context) {
	var book models.Book
	initializers.DB.Where("id = ?", c.Param("id")).Delete(&book)
	c.JSON(200, &book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	initializers.DB.Where("id = ?", c.Param("id")).First(&book)
	c.BindJSON(&book)
	initializers.DB.Save(&book)
	c.JSON(200, &book)

}
