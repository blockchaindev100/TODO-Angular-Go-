package routes

import (
	"fmt"
	"net/http"

	"github.com/blockchaindev100/todo/models"
	"github.com/blockchaindev100/todo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB) {
	service.SetDB(db)
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", getUsersRoute)
		userRoutes.GET("/:id", getUserByIdRoute)
		userRoutes.POST("/", createUserRoute)
		userRoutes.DELETE("/:id", deleteUserRoute)
	}
	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("/", getTodoRoute)
		todoRoutes.GET("/:id", getTodoByIdRoute)
		todoRoutes.POST("/", createTodoRoute)
		todoRoutes.DELETE("/:id", deleteTodoRoute)
		todoRoutes.PUT("/:id", updateTodoByIdRoute)
	}

}

func getUsersRoute(c *gin.Context) {
	data, err := service.GetUsers()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Error")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func getUserByIdRoute(c *gin.Context) {
	id := c.Param("id")
	data, err := service.GetUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "User Not Found")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func createUserRoute(c *gin.Context) {
	var data models.Users
	err := c.BindJSON(&data)
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

func deleteUserRoute(c *gin.Context) {
	err := service.DeleteUser(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Id not found")
		return
	}
	c.IndentedJSON(http.StatusOK, "Deleted successfully")
}

func getTodoRoute(c *gin.Context) {

}

func getTodoByIdRoute(c *gin.Context) {

}

func createTodoRoute(c *gin.Context) {

}

func deleteTodoRoute(c *gin.Context) {

}

func updateTodoByIdRoute(c *gin.Context) {

}
