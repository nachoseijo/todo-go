package controllers

import (
	"net/http"
	"todo-go/database"

	"github.com/gin-gonic/gin"
)

//GetAllTodos returns all todos from database
func GetAllTodos(context *gin.Context) {
	todos := database.FindAll()

	if todos == nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
	return
}

//CreateTodo inserts a todo in database
func CreateTodo(context *gin.Context) {
	var todo database.Todo
	context.BindJSON(&todo)

	insertedTodo, err := database.InsertOne(todo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
	}

	//needs header location
	context.JSON(http.StatusCreated, gin.H{
		"data": insertedTodo,
	})
	return
}

//GetSingleTodo returns a single todo from database
func GetSingleTodo(context *gin.Context) {
	todoID := context.Param("todoId")

	todo := database.FindOne(todoID)

	if todo == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
	return
}

//UpdateTodo updates a todo and saves in database
func UpdateTodo(context *gin.Context) {
	todoID := context.Param("todoId")
	todoInDB := database.FindOne(todoID)
	var todo database.Todo

	if todoInDB == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
		return
	}

	context.BindJSON(&todo)

	if database.UpdateOne(todoID, todo) != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
	return
}

//DeleteTodo deletes a todo from database
func DeleteTodo(context *gin.Context) {
	todoID := context.Param("todoId")

	if database.Delete(todoID) != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
	}

	context.JSON(http.StatusNoContent, gin.H{})
	return
}
