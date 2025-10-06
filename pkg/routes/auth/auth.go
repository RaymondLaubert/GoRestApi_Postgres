package auth

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	
	auth := route.Group("/auth") 
	{
		auth.GET("/login", login)
	}

}

// Function to Authenticate a User and Login
func login(context *gin.Context) {

}