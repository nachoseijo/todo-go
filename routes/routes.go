package routes

import (
	"net/http"

	"todo-go/controllers"

	"github.com/gin-gonic/gin"
)

//Routes creates API routes
func Routes(router *gin.Engine) {
	router.GET("/api/v1/todo", controllers.GetAllTodos)
	router.POST("/api/v1/todo", controllers.CreateTodo)
	router.GET("/api/v1/todo/:todoId", controllers.GetSingleTodo)
	router.PATCH("/api/v1/todo/:todoId", controllers.UpdateTodo)
	router.DELETE("/api/v1/todo/:todoId", controllers.DeleteTodo)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.String(http.StatusOK, `
		Welcome To API!
		Supported operations:
		- POST /todo
		- GET /todo
		- GET /todo/:todoId
		- PATCH /todo/:todoId
		- DELETE /todo/:todoId`,
	)
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Route Not Found",
	})
	return
}
