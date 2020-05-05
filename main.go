package main

import (
	"log"
	"os"

	"todo-go/database"
	"todo-go/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	apiPort := os.Getenv("API_PORT")
	databaseURI := os.Getenv("DATABASE_URI")
	databaseName := os.Getenv("DATABASE_NAME")
	collection := os.Getenv("COLLECTION")

	//Connect to Database
	database.Connect(collection, databaseName, databaseURI)

	//Inits a router
	router := gin.Default()
	//Enables cors for all origins
	router.Use(cors.Default())
	//Router Handlers
	routes.Routes(router)

	log.Fatal(router.Run(":" + apiPort))
}
