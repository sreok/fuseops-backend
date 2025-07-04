package model

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("table_prefix", "t_") // 只设置一次默认值
}

// User 用户表结构体
type User struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"user_id"` // 用户ID
	UserName string `gorm:"column:username" json:"username"`     // 用户名
	Password string `gorm:"column:password" json:"password"`     // 密码
	RealName string `gorm:"column:realname" json:"realname"`     // 真实姓名
	Email    string `gorm:"column:email" json:"email"`           // 邮箱
	Phone    string `gorm:"column:phone" json:"phone"`           // 手机号
	Avatar   string `gorm:"column:avatar" json:"avatar"`         // 头像地址
	RoleID   int    `gorm:"column:role_id" json:"role_id"`       // 角色ID
	Status   int    `gorm:"column:status" json:"status"`         // 状态 1正常0禁用
}

// TableName 设置表名
func (User) TableName() string {
	return viper.GetString("table_prefix") + "user"
}

// Role 规则表结构体
type Role struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"role_id"` // 用户ID
	RoleName string `gorm:"column:role_name" json:"role_name"`   // 角色名称
	RoleCode string `gorm:"column:role_code" json:"role_code"`   // 角色编码
}

// TableName 设置表名
func (Role) TableName() string {
	return viper.GetString("table_prefix") + "role"
}

// User_Role 规则表结构体
type UserRole struct {
	ID     uint `gorm:"primaryKey;column:id" json:"id"` // 用户ID
	UserID uint `gorm:"column:user_id" json:"user_id"`  // 角色名称
	RoleID uint `gorm:"column:role_id" json:"role_id"`  // 角色编码
}

// TableName 设置表名
func (UserRole) TableName() string {
	return viper.GetString("table_prefix") + "user_role"
}
