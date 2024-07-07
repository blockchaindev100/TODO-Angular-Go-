package routes

import (
	"github.com/blockchaindev100/todo/handler"
	"github.com/blockchaindev100/todo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB) {
	service.SetDB(db)
	userRoutes := router.Group("/users")
	userRoutes.GET("/", handler.GetUsersRoute)
	userRoutes.GET("/:id", handler.GetUserByIdRoute)
	userRoutes.POST("/", handler.CreateUserRoute)
	userRoutes.DELETE("/:id", handler.DeleteUserRoute)
	todoRoutes := router.Group("/todo")
	todoRoutes.GET("/", handler.GetTodoRoute)
	todoRoutes.GET("/:id", handler.GetTodoByIdRoute)
	todoRoutes.POST("/", handler.CreateTodoRoute)
	todoRoutes.DELETE("/:id", handler.DeleteTodoRoute)
	todoRoutes.PUT("/:id", handler.UpdateTodoByIdRoute)

}
