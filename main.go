package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/madbull12/gin_postgres/initializers"
	"github.com/madbull12/gin_postgres/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connect()
	initializers.SyncToDatabase()
}

func main() {
	router := gin.New()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello world")
	// })
	// initializers.Connect()
	fmt.Println("1")
	routes.BookRoute(router)
	routes.UserRoute(router)

	router.Run(":8080")
}
