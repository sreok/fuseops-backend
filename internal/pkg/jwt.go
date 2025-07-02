package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// GenerateToken 生成 JWT
func GenerateToken(userID uint, userName string) (string, error) {
	jwtSecret := []byte(viper.GetString("jwt.secret"))
	expireHours := viper.GetInt("jwt.expire_hours")
	claims := jwt.MapClaims{
		"user_id":   userID,
		"user_name": userName,
		"exp":       time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	jwtSecret := []byte(viper.GetString("jwt.secret"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
