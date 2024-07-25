package handler

import (
	"fmt"
	"net/http"

	"github.com/blockchaindev100/todo/models"
	"github.com/blockchaindev100/todo/repository"
	"github.com/blockchaindev100/todo/service"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	data, err := repository.GetUsers()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Error")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	data, err := repository.GetUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "User Not Found")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func CreateUser(c *gin.Context) {
	var data models.Users
	err := c.BindJSON(&data)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error: internal error")
		return
	}
	errInCreateUser := repository.CreateUsers(&data)
	if errInCreateUser != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Can't create User")
		return
	}
	service.SendMail(data.Email)
	c.IndentedJSON(http.StatusOK, "User created successfully")
}

func DeleteUser(c *gin.Context) {
	err := repository.DeleteUser(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusOK, "Deleted successfully")
}
