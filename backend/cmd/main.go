package main

import (
	"log"

	"github.com/blockchaindev100/todo/adaptor"
	"github.com/blockchaindev100/todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := adaptor.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetUpRoutes(router, db)

	err = router.Run("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

}
