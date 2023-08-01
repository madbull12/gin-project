package main

import (
	"github.com/gin-gonic/gin"
	"github.com/madbull12/gin_postgres/initializers"
	"github.com/madbull12/gin_postgres/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main() {
	router := gin.New()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello world")
	// })
	// initializers.Connect()
	routes.BookRoute(router)

	router.Run(":8080")
}
