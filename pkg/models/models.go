package models

import (
	"time"
)

type User struct {
	id int64
	username string
	password string
}

type TodoList struct {
	title string
	description string
	due time.Time
	completed time.Time
}