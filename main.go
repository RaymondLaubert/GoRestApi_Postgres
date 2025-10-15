package main

import (
	"fmt"
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
		fmt.Printf("Unable to Establish Connection with the Database: %s", err.Error())
		os.Exit(1)
	}
	
	// Create the Database Tables
	err = dbConn.CreateDatabaseTables()
	if err != nil {
		fmt.Printf("Unable to Create Database Tables: %s", err.Error())
		os.Exit(2)
	}

	// Create a Gin Router with Default Middleware (Logger and Recovery)
	router := gin.Default()

	// Add All Authentication Routes
	authRouter := auth.AuthRouter {Router: router, Db: &dbConn}
	authRouter.Routes()

	// Add All User Routes
	userRouter := users.UserRouter {Router: router, Db: &dbConn}
	userRouter.Routes()

	// Add All ToDo Routes
	todoRouter := todo.TodoRouter {Router: router, Db: &dbConn}
	todoRouter.Routes()

	// Start Server on Port 8080 (Default)
	router.Run()

}