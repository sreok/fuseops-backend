package model

import (
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("table_prefix", "t_") // 只设置一次默认值
}

// User 用户表结构体
type User struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`       // 用户ID
	UserName   string    `gorm:"column:user_name" json:"username"`     // 用户名
	RealName   string    `gorm:"column:real_name" json:"realname"`     // 真实姓名
	Password   string    `gorm:"column:password" json:"password"`      // 密码
	Avatar     string    `gorm:"column:avatar" json:"avatar"`          // 头像地址
	RoleID     int       `gorm:"column:role_id" json:"roleid"`         // 角色ID
	GradeID    int       `gorm:"column:grade_id" json:"gradeid"`       // 班级ID
	CreateTime time.Time `gorm:"column:create_time" json:"createtime"` // 创建时间
	Status     int       `gorm:"column:status" json:"status"`          // 状态 1正常0禁用
	IsDeleted  int       `gorm:"column:is_deleted" json:"isdeleted"`   // 逻辑删除 0未删除 1已删除
}

// TableName 设置表名
func (User) TableName() string {
	return viper.GetString("table_prefix") + "user"
}
