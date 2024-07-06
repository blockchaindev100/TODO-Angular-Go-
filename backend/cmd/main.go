package main

import (
	"log"

	"github.com/blockchaindev100/todo/models"
	"github.com/blockchaindev100/todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetUpRoutes(router, db)

	router.Run("localhost:8000")

}
