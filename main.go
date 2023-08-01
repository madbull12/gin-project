package main

import (
	"github.com/gin-gonic/gin"
	"github.com/madbull12/gin_postgres/config"
	"github.com/madbull12/gin_postgres/routes"
)

func main() {
	router := gin.New()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello world")
	// })
	config.Connect()
	routes.UserRoute(router)

	router.Run(":8080")
}
