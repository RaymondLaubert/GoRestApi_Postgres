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

func EstablishDatabaseConnection(connString string) (Database, error) {
	
	dbConn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		errorMessage := fmt.Errorf("Unable to Connect to Database: %v\n", err)
		return Database{}, errorMessage
	}

	return Database {conn: dbConn}, nil
	
}

func (db *Database) GetUser(id int64) (models.User, error) {
	row := db.conn.QueryRow(context.Background(), "SELECT * FROM User WHERE id=$1")
}

func (db *Database) GetAllUsers() ([]models.User, error) {
	row := db.conn.QueryRow(context.Background(), "SELECT * FROM User")
}

func (db *Database) CreateUser() (models.User, error) {

}

func (db *Database) GetTodoList() (models.TodoList, error) {

}

