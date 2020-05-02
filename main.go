package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nachoseijo/todo/routes"
)

func main() {
	//init a router
	router := gin.Default()

	Routes.routes(router)

	log.Fatal(router.Run(":8090"))
}
