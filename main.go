package main

func main() {
	
		r := gin.Default()
	    r.POST("/signup",controllers.Signup)
	
	
		r.Run(":3030")
	}
	
	
