package users

import (
	"net/http"
	"strconv"

	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/database"
	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/models"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Router *gin.Engine
	Db *database.Database
}

// Function Adds All the User Routes to the Router
func (ur *UserRouter) Routes() {

	users := ur.Router.Group("/users")
	{
		users.GET("/user", ur.getUser)
		users.GET("/allUsers", ur.getAllUsers)
		users.POST("/createUser", ur.createUser)
		users.PUT("/updateUser", ur.updateUser)
		users.DELETE("/deleteUser", ur.deleteUser)
	}

}

// Function to Get a User
func (ur *UserRouter) getUser(context *gin.Context) {
	
	// Obtain the User ID from the Parameters in Path
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.String(http.StatusNoContent, "User ID Missing: %w", err)
		return
	}

	// Obtain the User
	user, err := ur.Db.GetUser(userId)
	if err != nil {
		context.String(http.StatusNotFound, "User Not Found: %w", err)
		return
	}

	// Return a JSON Response
	context.JSON(http.StatusOK, gin.H {
		"message": user,
	})

}

// Function to Get All Users
func (ur *UserRouter) getAllUsers(context *gin.Context) {

	// Obtain the Users
	users, err := ur.Db.GetAllUsers()
	if err != nil {
		context.String(http.StatusNotFound, "%w", err)
		return
	}

	// Return a JSON Response
	context.JSON(http.StatusOK, gin.H {
		"message": users,
	})

}

// Function to Create a User
func (ur *UserRouter) createUser(context *gin.Context) {

	// Obtain the Username and Password

	// Create the User via the User Model
	user := models.User {
		Username: "",
		Password: "",
	}

	// Return a JSON Response with the User
	context.JSON(http.StatusOK, gin.H {
		"message": user,
	})

}

// Function to Update a User
func (ur *UserRouter) updateUser(context *gin.Context) {

	// Obtain the New Username and/or Password
	

	// Create the User via the User Model
	user := models.User {
		Id: 1,
		Username: "",
		Password: "",
	}

	// Return a JSON Response with the User
	context.JSON(http.StatusOK, gin.H {
		"message": user,
	})

}

// Function to Delete a User
func (ur *UserRouter) deleteUser(context *gin.Context) {

	// Obtain the User via Binding
	user := models.User{}

	// Attempt to Delete the User
	err := ur.Db.DeleteUser(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	// Return a Successful Response
	context.JSON(http.StatusOK, gin.H {
		"message": "Successfully Deleted User.",
	})

}