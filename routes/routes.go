package routes

import (
	"go-jwt-api/controllers"
	"go-jwt-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/protected", controllers.ProtectedEndpoint)

	return r
}
