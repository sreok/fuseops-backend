package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// GenerateToken 生成 JWT
// userID: 用户ID
// userName: 用户名
// 返回生成的 token 字符串和错误信息
func GenerateToken(userID uint, userName string) (string, error) {
	jwtSecret := []byte(viper.GetString("jwt.secret")) // 获取 JWT 密钥
	expireHours := viper.GetInt("jwt.expire_hours")    // 获取过期时间（小时）
	claims := jwt.MapClaims{
		"user_id":   userID,                                                        // 用户ID
		"user_name": userName,                                                      // 用户名
		"exp":       time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(), // 过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 创建带有声明的 token
	return token.SignedString(jwtSecret)                       // 签名并返回 token 字符串
}

// ParseToken 解析 JWT
// tokenString: 待解析的 token 字符串
// 返回解析后的 claims 和错误信息
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	jwtSecret := []byte(viper.GetString("jwt.secret")) // 获取 JWT 密钥
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil // 返回密钥用于校验
	})
	if err != nil || !token.Valid { // 校验失败或 token 无效
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil // 返回 claims
}
