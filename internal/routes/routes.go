package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/lms_backend/internal/app/handler"
	"github.com/irawankilmer/lms_backend/internal/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	api := r.Group("/api")
	api.Use(middleware.ClientAuth())

	api.POST("/roles", handler.CreateRole)
	api.GET("/roles", handler.GetRoles)
	api.GET("/roles/:id", handler.GetRole)
	api.PUT("/roles/:id", handler.UpdateRole)
	api.DELETE("/roles/:id", handler.DeleteRole)

	// WebSocket
	r.GET("/ws", handler.WebSocketHandler)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
