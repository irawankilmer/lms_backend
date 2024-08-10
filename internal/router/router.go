package router

import (
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/config"
	_ "github.com/irawankilmer/lms_backend/docs" // Swagger docs
	"github.com/irawankilmer/lms_backend/internal/handler"
	"github.com/irawankilmer/lms_backend/internal/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// SetupRouter initializes the routes for the application
func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.POST("/login", handler.Login)

	// Protected routes
	protected := r.Group("/users", middleware.AuthMiddleware(cfg))
	{
		protected.POST("/", handler.CreateUser)
		protected.GET("/", handler.GetAllUsers)
		protected.GET("/:id", handler.GetUserByID)
		protected.PUT("/:id", handler.UpdateUser)
		protected.DELETE("/:id", handler.DeleteUser)
	}

	return r
}
