package users

import "github.com/gin-gonic/gin"


// Function Adds All the User Routes to the Router
func Routes(route *gin.Engine) {

	users := route.Group("/users")
	{
		users.GET("/user", getUser)
		users.GET("/allUsers", getAllUsers)
		users.POST("/createUser", createUser)
		users.PUT("/updateUser", updateUser)
		users.DELETE("/deleteUser", deleteUser)
	}

}

// Function to Get a User
func getUser(context *gin.Context) {

}

// Function to Get All Users
func getAllUsers(context *gin.Context) {

}

// Function to Create a User
func createUser(context *gin.Context) {

}

// Function to Update a User
func updateUser(context *gin.Context) {

}

// Function to Delete a User
func deleteUser(context *gin.Context) {

}