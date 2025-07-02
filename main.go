// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"fmt"
	"log"

	"fuseops-backend/internal/api"
	"fuseops-backend/internal/repository"

	_ "fuseops-backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FuseOps API
// @version 1.0.0
// @description 基于client-go和gin框架后端api服务
// @termsOfService
func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件失败: %v\n", err)
	}
}

func main() {
	initConfig()
	repository.InitDB()

	r := gin.Default()

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := r.Group("/api")
	{
		api.RegisterAuthRoutes(apiGroup)
		api.RegisterUserRoutes(apiGroup)
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8002"
	}
	err := r.Run(":" + port)
	if err != nil {
		log.Println(err)
		return
	}
}
