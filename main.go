package main

import (
	"test/repo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/music", controllers.Music)
	r.GET("/search", controllers.Search)
	r.GET("/allmusic", controllers.AllMusics)

	r.Run(":3030")
}
