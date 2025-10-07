package models

import (
	"time"
)

type User struct {
	Id 			int64	`json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
}

type TodoList struct {
	Title 		string		`json:"Title"`
	Description string		`json:"Description"`
	Due 		time.Time	`json:"Due"`
	Completed 	bool		`json:"Completed"`
}