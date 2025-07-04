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
		userGroup.POST("", CreateUser)
		userGroup.GET("", GetUser)
		userGroup.PUT("", UpdateUser)
		userGroup.DELETE("", DeleteUser)
	}
}

// RegisterRoleRoutes 注册角色相关的路由
func RegisterRoleRoutes(api *gin.RouterGroup) {
	roleGroup := api.Group("/roles")
	roleGroup.Use(middleware.JWTAuth())
	{
		roleGroup.POST("", CreateRole)
		roleGroup.GET("", GetRole)
		roleGroup.PUT("", UpdateRole)
		roleGroup.DELETE("", DeleteRole)
	}
	{
		roleGroup.GET("/bind", GetUserBindRole)
	}
}

// RegisterAuthRoutes 注册认证相关路由
func RegisterAuthRoutes(api *gin.RouterGroup) {
	api.POST("/login", handler.Login)
}
