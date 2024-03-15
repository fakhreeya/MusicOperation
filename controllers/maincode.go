package controllers

func Signup(c *gin.Context)  {
	var SignUpTemp structs.SignUpStruct
	c.ShouldBindJSON(&SignUpTemp)

	if SignUpTemp.Name == "" || SignUpTemp.Surname == "" || SignUpTemp.Login == "" || SignUpTemp.Password == "" {
		c.JSON(404, "Error")
	} else {
}
}






























