package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB) {
	routes := router.Group("/users")
	{
		routes.GET("/", getUserRoute)
		routes.POST("/")
		routes.DELETE("/")
	}
}

func getUserRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello")
}
