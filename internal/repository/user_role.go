package repository

import (
	"fmt"
	"fuseops-backend/internal/model"
)

// GetRolesByUserID 根据用户ID查找所有角色
func GetRoleIDsByUserID(userID uint) ([]uint, error) {
	var roleIDs []uint
	var userRoles []model.UserRole

	// 从配置文件获取表名
	tableName := model.UserRole{}.TableName()

	if err := DB.Table(tableName).Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		fmt.Println("查询角色ID失败:", err)
		return roleIDs, err
	}
	for _, ur := range userRoles {
		roleIDs = append(roleIDs, ur.RoleID)
	}
	return roleIDs, nil
}

// GetRoleCodeByRoleIDs 根据角色ID查找所有角色代码
func GetRoleCodeByRoleIDs(roleIDs []uint) ([]string, error) {

	var Roles []model.Role
	var RoleCodes []string

	// 从配置文件获取表名
	tableName := model.Role{}.TableName()

	for _, id := range roleIDs {

		if err := DB.Table(tableName).Where("id = ?", id).Find(&Roles).Error; err != nil {
			fmt.Println("查询角色Code失败:", err)
			return nil, err
		}
		// You may need to extract RoleCode from Roles, not ur
		for _, role := range Roles {
			RoleCodes = append(RoleCodes, role.RoleCode)
		}
	}
	return RoleCodes, nil
}
