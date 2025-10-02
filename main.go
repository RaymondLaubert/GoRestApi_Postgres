package main

import (
	"fmt"
	"net/http"

	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	// Create the Database URL Connection String
	databaseUrl := "postgres://xyeles:s3cret@localhost:5432/goapi_postgres_db"

	// Attempt to Establish a Connection with the PostgresDB
	dbConn, err := database.EstablishDatabaseConnection(databaseUrl)
	if err != nil {
		fmt.Printf("Unable to Establish Connection with the Database: %w", err)
	}
	
	// Create a Gin Router with Default Middleware (Logger and Recovery)
	router := gin.Default()

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