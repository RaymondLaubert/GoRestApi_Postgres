package todo

import (
	"github.com/RaymondLaubert/GoRestApi_Postgres/pkg/database"

	"github.com/gin-gonic/gin"
)

type TodoRouter struct {
	Router *gin.Engine
	Db *database.Database
}

func (tr *TodoRouter) Routes() {
	
	todo := tr.Router.Group("/todo")
	{
		todo.GET("/todoList", tr.getTodoList)
	}

}

// Function to Get the ToDo List
func (tr *TodoRouter) getTodoList(context *gin.Context) {

}