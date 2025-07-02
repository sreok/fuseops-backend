package api

import (
	"fuseops-backend/internal/handler"
	"fuseops-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(api *gin.RouterGroup) {
	userGroup := api.Group("/users")
	userGroup.Use(middleware.JWTAuth())
	{
		userGroup.POST("", handler.CreateUser)
		userGroup.GET("", handler.GetUser)
		userGroup.PUT("", handler.UpdateUser)
		userGroup.DELETE("", handler.DeleteUser)
	}
}

// RegisterAuthRoutes 注册认证相关路由
func RegisterAuthRoutes(api *gin.RouterGroup) {
	api.POST("/login", handler.Login)
}
