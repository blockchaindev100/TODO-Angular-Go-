package handler

import (
	"net/http"

	"github.com/blockchaindev100/todo/models"
	"github.com/blockchaindev100/todo/repository"
	"github.com/gin-gonic/gin"
)

func GetTodoRoute(c *gin.Context) {
	var user models.Users
	c.Bind(&user)
	data, err := repository.GetTodo(user.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Can't get todo")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func GetTodoByIdRoute(c *gin.Context) {
	var user models.Users
	c.Bind(&user)
	id := c.Param("id")
	data, err := repository.GetTodoByID(user.Id, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Can't get todo")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func CreateTodoRoute(c *gin.Context) {
	var todo models.Todo
	c.Bind(&todo)
	err := repository.CreateTodo(&todo)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Todo not created")
		return
	}
	c.IndentedJSON(http.StatusOK, "Todo created successfully")
}

func DeleteTodoRoute(c *gin.Context) {

}

func UpdateTodoByIdRoute(c *gin.Context) {

}
