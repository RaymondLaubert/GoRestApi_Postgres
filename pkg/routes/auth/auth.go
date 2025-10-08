package auth

import (
	"net/http"

	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/database"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type LoginDetails struct {
	Username	string	`json:"username" binding:"required"`
	Password 	string	`json:"password" binding:"required"`
}

type AuthRouter struct {
	Router *gin.Engine
	Db *database.Database
}

func (ar *AuthRouter) Routes() {
	
	auth := ar.Router.Group("/auth") 
	{
		auth.POST("/login", ar.login)
	}

}

// Function to Authenticate a User and Login
func (ar *AuthRouter) login(context *gin.Context) {

	var details LoginDetails

	// Attempt to Bind the Login Information
	if err := context.ShouldBindJSON(&details); err != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	// Attempt to Authenticate the User with Login Credentials
	userFound, err := ar.Db.AuthenticateUser(details.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error": "User Not Found.",
		})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(details.Password)); err != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid Password.",
		})
		return
	}

	// Return Login Authorized
	context.JSON(http.StatusOK, gin.H {
		"status": "Login Authorized",
	})

}

// Function to Check the Authentication of a User
func (ar *AuthRouter) CheckAuth(context *gin.Context) {



}