package todo

import (
	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/database"

	"github.com/gin-gonic/gin"
)

type TodoRoutes struct {
	router *gin.Engine
	db *database.Database
}

func (tr *TodoRoutes) Routes() {
	
	todo := tr.router.Group("/todo")
	{
		todo.GET("/todoList", tr.getTodoList)
	}

}

// Function to Get the ToDo List
func (tr *TodoRoutes) getTodoList(context *gin.Context) {

}