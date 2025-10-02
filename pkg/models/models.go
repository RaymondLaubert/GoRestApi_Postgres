package models

import (
	"time"
)

type User struct {
	Id int64
	Username string
	Password string
}

type TodoList struct {
	Title string
	Description string
	Due time.Time
	Completed time.Time
}