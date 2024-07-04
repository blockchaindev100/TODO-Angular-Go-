package main

import (
	"log"

	"github.com/blockchaindev100/todo/models"
	"github.com/gin-gonic/gin"
)

type Roles struct {
	Id        int
	Role_type string
}

type Todo struct {
	Id        int
	Task_name string
	Status    string
	Active    bool
	User_id   int
}

type Users struct {
	Id                int
	Name              string
	Email             string
	Roleid            int
	Verification_code string
	Is_verified       bool
}

func main() {
	db, err := models.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

}
