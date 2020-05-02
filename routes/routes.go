package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Routes creates API routes
func Routes(router *gin.Engine) {
	router.GET("/", helloHandler)
	router.GET("/one", endpointOne)
	router.GET("/two", endpointTwo)
}

func helloHandler(context *gin.Context) {
	context.String(http.StatusOK, "Hello world with GIN!")
}

func endpointOne(context *gin.Context) {
	context.String(http.StatusOK, "Accessing endpoint One with Gin")
}

func endpointTwo(context *gin.Context) {
	context.String(http.StatusOK, "Accessing endpoint Two with Gin")
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
