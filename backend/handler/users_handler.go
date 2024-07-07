package handler

import (
	"fmt"
	"net/http"

	"github.com/blockchaindev100/todo/models"
	"github.com/blockchaindev100/todo/service"
	"github.com/gin-gonic/gin"
)

func GetUsersRoute(c *gin.Context) {
	data, err := service.GetUsers()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Error")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func GetUserByIdRoute(c *gin.Context) {
	id := c.Param("id")
	data, err := service.GetUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "User Not Found")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func CreateUserRoute(c *gin.Context) {
	var data models.Users
	err := c.BindJSON(&data)
	fmt.Println(err)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error: internal error")
		return
	}
	errInCreateUser := service.CreateUsers(&data)
	if errInCreateUser != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Can't create User")
		return
	}
	service.SendMail(data.Email)
	c.IndentedJSON(http.StatusOK, "User created successfully")
}

func DeleteUserRoute(c *gin.Context) {
	err := service.DeleteUser(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Id not found")
		return
	}
	c.IndentedJSON(http.StatusOK, "Deleted successfully")
}
