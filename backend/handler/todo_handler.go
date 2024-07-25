package handler

import (
	"fmt"
	"net/http"

	"github.com/blockchaindev100/todo/models"
	"github.com/blockchaindev100/todo/repository"
	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	var user models.Users
	c.BindJSON(&user)
	fmt.Println(user.Id)
	data, err := repository.GetTodo(user.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Can't get todo")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func GetTodoById(c *gin.Context) {
	var user models.Users
	c.BindJSON(&user)
	id := c.Param("id")
	data, err := repository.GetTodoByID(user.Id, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Can't get todo")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := repository.CreateTodo(&todo)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Todo not created")
		return
	}
	c.IndentedJSON(http.StatusOK, "Todo created successfully")
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteTodo(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, "Todo deleted successfully")
}

func UpdateTodoById(c *gin.Context) {
	var data models.Todo
	c.BindJSON(&data)
	err := repository.UpdateTodo(&data)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Todo not updated")
	}
	c.IndentedJSON(http.StatusOK, data)
}
