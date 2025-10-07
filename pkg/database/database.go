package database

import (
	"context"
	"fmt"

	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
}

// Function to Establish a Connection to the Database
func EstablishDatabaseConnection(connString string) (Database, error) {
	
	dbConn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		errorMessage := fmt.Errorf("Unable to Connect to Database: %v\n", err)
		return Database{}, errorMessage
	}

	return Database {Conn: dbConn}, nil
	
}

func (db *Database) CreateDatabaseTables() (error) {

	// Start the Transaction
	transaction, err := db.Conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to Begin Transaction (CreateDatabaseTables): %w", err)
	}

	// Ensure the Transaction will be Rolled Back if not Committed
	defer transaction.Rollback(context.Background())
	
	// Create the Query to Create the Users Table
	usersQuery := `CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(100) NOT NULL
	);`

	// Create the Query to Create the ToDo List Table
	todoListQuery := `CREATE TABLE todo (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description VARCHAR(255),
		due TIMESTAMP,
		completed BOOLEAN
	);`

	// Execute the Create Users Table Query
	_, err = transaction.Exec(context.Background(), usersQuery)
	if err != nil {
		return fmt.Errorf("Unable to Create the Users Table: %w", err)
	}

	// Execute the Create ToDo List Table Query
	_, err = transaction.Exec(context.Background(), todoListQuery)
	if err != nil {
		return fmt.Errorf("Unable to Create the ToDo List Table: %w", err)
	}

	// Commit the Transaction
	if err = transaction.Commit(context.Background()); err != nil {
		return fmt.Errorf("Error Committing Transaction (CreateDatabaseTables): %w", err)
	}

	return nil

}

// Function to Help Authenticate a User
func (db *Database) AuthenticateUser(username string) (models.User, error) {

	// Create the User to Return
	user := models.User{}

	// Query the Database for the User
	err := db.Conn.QueryRow(context.Background(), "SELECT * FROM users WHERE username=$1", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("User Not Found: %w", err)
	}

	// Return the Found User
	return user, nil

}

// Function to Get a User from the Database
func (db *Database) GetUser(id int64) (models.User, error) {
	
	// Create the User to Return
	user := models.User{}

	// Query the Database for the User and Scan the Row Received for the User ID, Username, and Password
	err := db.Conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("Unable to Find User: %w", err)
	}

	// Return the Found User
	return user, nil

}

// Function to Get All the Users from the Database in a List
func (db *Database) GetAllUsers() ([]models.User, error) {
	
	// Query the Database for All Users
	rows, err := db.Conn.Query(context.Background(), "SELECT * FROM users")
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
	transaction, err := db.Conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to Begin Transaction (CreateUser): %w", err)
	}

	// Ensure the Username doesn't Exist Already
	existingUser := models.User{}
	err = db.Conn.QueryRow(context.Background(), "SELECT * FROM users WHERE username=$1", user.Username).Scan(&existingUser.Id, &existingUser.Username, &existingUser.Password)
	if err == nil {
		return fmt.Errorf("Username Already Exists.")
	}

	// Create the Query
	query := `INSERT INTO users (username, password) VALUES (@userName, @userPassword)`

	// Create the Password Hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Unable to Hash Password: %w", err)
	}

	// Create the Named Arguments
	args := pgx.NamedArgs{
		"userName": user.Username,
		"userPassword": string(passwordHash),
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
		return fmt.Errorf("Error Committing Transaction (CreateUser): %w", err)
	}

	return nil

}

// Function to Update a User
func (db *Database) UpdateUser(user models.User) (error) {

	return nil

}

// Function to Delete a User from the Database
func (db *Database) DeleteUser(user models.User) (error) {
	
	// Start the Transaction by Calling Begin
	transaction, err := db.Conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to Begin Transaction (DeleteUser): %w", err)
	}

	// Ensure the Transaction will be Rolled Back if not Committed
	defer transaction.Rollback(context.Background())
	
	// Create the Delete Statement
	query := `DELETE FROM users WHERE id = $1`

	// Execute the Deletion of the User
	_, err = transaction.Exec(context.Background(), query, user.Id)
	if err != nil {
		return fmt.Errorf("Unable to Delete User: %w", err)
	}

	// Commit the Transaction
	if err = transaction.Commit(context.Background()); err != nil {
		return fmt.Errorf("Error Committing Transaction (DeleteUser): %w", err)
	}

	return nil

}

// Function to Get the User's ToDo List from the Database
func (db *Database) GetTodoList(userId int64) ([]models.TodoList, error) {

	// Create the Query
	query := `SELECT * FROM todo WHERE id = $1`

	// Query the Database for the User's ToDo List
	rows, err := db.Conn.Query(context.Background(), query, userId)
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

