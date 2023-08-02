package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madbull12/gin_postgres/controller"
)

func UserRoute(router *gin.Engine) {
	// router.GET("/", controller.GetBooks)
	// router.GET("/:id", controller.GetSingleUser)
	router.POST("/signup", controller.Signup)
	router.POST("/signin", controller.Signin)
	// router.DELETE("/:id", controller.DeleteBook)
	// router.PUT("/:id", controller.UpdateBook)

}
