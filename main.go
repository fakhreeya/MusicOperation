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
	
	func Cors(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.43.246:5500")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
	
		c.Next()
	}
	
