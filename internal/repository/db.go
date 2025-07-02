package repository

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 构建 MySQL DSN（数据源名称）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.user"),     // 数据库用户名
		viper.GetString("db.password"), // 数据库密码
		viper.GetString("db.host"),     // 数据库主机地址
		viper.GetString("db.port"),     // 数据库端口
		viper.GetString("db.dbname"),   // 数据库名称
	)

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
}
