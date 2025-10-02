package database

import (
	"context"
	"fmt"

	"github.com/RaymondLaubert/GoRestApiPostgres/pkg/models"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgx.Conn
}

// Function to Establish a Connection to the Database
func EstablishDatabaseConnection(connString string) (Database, error) {
	
	dbConn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		errorMessage := fmt.Errorf("Unable to Connect to Database: %v\n", err)
		return Database{}, errorMessage
	}

	return Database {conn: dbConn}, nil
	
}

// Function to Get a User from the Database
func (db *Database) GetUser(id int64) (models.User, error) {
	
	// Create the User to Return
	user := models.User{}

	// Query the Database for the User and Scan the Row Received for the User ID, Username, and Password
	err := db.conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("Unable to Find User: %w", err)
	}

	return user, nil

}

// Function to Get All the Users from the Database in a List
func (db *Database) GetAllUsers() ([]models.User, error) {
	
	// Query the Database for All Users
	rows, err := db.conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("Unable to Query Users: %w", err)
	}

	// Defer the Closure of the Rows Received
	// Ensures the Rows are Closed even if an Error Occurs Later
	defer rows.Close()

	// Obtain the Users from the Rows Received
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, fmt.Errorf("Unable to Collect Users from the Rows: %w", err)
	} 

	return users, nil

}

// Function to Create a User and Insert them into the Database
func (db *Database) CreateUser(user models.User) (error) {

	// Start the Transaction by Calling Begin
	transaction, err := db.conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to Begin Transaction: %w", err)
	}

	// Create the Query
	query := `INSERT INTO users (username, password) VALUES (@userName, @userPassword)`

	// Create the Named Arguments
	args := pgx.NamedArgs{
		"userName": user.Username,
		"userPassword": user.Password,
	}

	// Ensure the Transaction will be Rolled Back if not Committed
	defer transaction.Rollback(context.Background())

	// Execute the Insertion of the User
	_, err = transaction.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Unable to Insert User: %w", err)
	}

	// Commit the Transaction
	if err = transaction.Commit(context.Background()); err != nil {
		return fmt.Errorf("Error Committing Transaction: %w", err)
	}

	return nil

}

// Function to Get the User's ToDo List from the Database
func (db *Database) GetTodoList(userId int64) ([]models.TodoList, error) {

	// Create the Query
	query := `SELECT * FROM todo WHERE id=$1`

	// Query the Database for the User's ToDo List
	rows, err := db.conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil, fmt.Errorf("Unable to Query ToDo List: %w", err)
	}

	// Defer the Closure of the Rows Received
	// Ensures the Rows are Closed even if an Error Occurs Later
	defer rows.Close()

	// Obtain the ToDo List from the Rows Received
	todoList, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TodoList])
	if err != nil {
		return nil, fmt.Errorf("Unable to Collect ToDo List from the Rows: %w", err)
	}

	return todoList, nil

}

