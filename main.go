package main

import (
	"log"

	"todo-go/database"
	"todo-go/routes"

	"github.com/gin-gonic/gin"
)

//Needs to be modified in order to take the value from env vars
const databaseURI = "mongodb://database:27017"
const databaseName = "todo-go"
const collection = "todos"

func main() {

	//Connect to Database
	database.Connect(collection, databaseName, databaseURI)

	//Inits a router
	router := gin.Default()

	//Router Handlers
	routes.Routes(router)

	log.Fatal(router.Run(":80"))
}
