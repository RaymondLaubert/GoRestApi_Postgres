package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/database"
	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/routes/auth"
	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/routes/todo"
	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/routes/users"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create the Database URL Connection String
	databaseUrl := "postgres://xyeles:s3cret@localhost:5432/goapi_postgres_db"

	// Attempt to Establish a Connection with the PostgresDB
	dbConn, err := database.EstablishDatabaseConnection(databaseUrl)
	if err != nil {
		fmt.Printf("Unable to Establish Connection with the Database: %s", err)
		os.Exit(1)
	}
	
	// Create the Database Tables
	err = dbConn.CreateDatabaseTables()
	if err != nil {
		os.Exit(2)
	}

	// Create a Gin Router with Default Middleware (Logger and Recovery)
	router := gin.Default()

	// Add All Authentication Routes
	auth.Routes(router)

	// Add All User Routes
	users.Routes(router)

	// Add All ToDo Routes
	todo.Routes(router)

	// Define a Simple GET Endpoint
	router.GET("/ping", func(context *gin.Context) {
		// Return JSON Response
		context.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})

	// Start Server on Port 8080 (Default)
	router.Run()

}