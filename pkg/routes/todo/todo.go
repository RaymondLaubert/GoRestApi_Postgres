package todo

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	
	todo := route.Group("/todo")
	{
		todo.GET("/todoList", GetTodoList)
	}

}

// Function to Get the ToDo List
func GetTodoList(context *gin.Context) {

}