package routes

import (
	"github.com/blockchaindev100/todo/config"
	"github.com/blockchaindev100/todo/handler"
	"github.com/blockchaindev100/todo/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB) {
	repository.SetDB(db)
	config.InitConfig()
	router.Use(handler.CorsMiddleware())
	router.GET("/google_login", handler.GoogleLogin)
	// router.GET("/google_callback", handler.GoogleCallback)
	router.POST("/api/google-login", handler.GoogleLoginHandler)
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", handler.GetUsers)
		userRoutes.GET("/:id", handler.GetUserById)
		userRoutes.POST("/", handler.CreateUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
	}
	todoRoutes := router.Group("/todo")
	{
		todoRoutes.GET("/", handler.GetTodo)
		todoRoutes.GET("/:id", handler.GetTodoById)
		todoRoutes.POST("/", handler.CreateTodo)
		todoRoutes.DELETE("/:id", handler.DeleteTodo)
		todoRoutes.PUT("/:id", handler.UpdateTodoById)
	}
}
