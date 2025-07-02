package middleware

import (
	"fuseops-backend/internal/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 校验 JWT Token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		// 检查 Authorization 是否存在且以 Bearer 开头
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供有效的token"})
			c.Abort()
			return
		}
		// 去除 Bearer 前缀，获取 token 字符串
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// 解析并校验 token
		claims, err := pkg.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效或已过期"})
			c.Abort()
			return
		}
		// 将解析出的 claims 信息保存到上下文
		c.Set("claims", claims)
		// 继续后续处理
		c.Next()
	}
}
