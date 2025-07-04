package api

import (
	"fuseops-backend/internal/model"
	"fuseops-backend/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRole 新建角色
// @Summary 新建角色
// @Description 新增一个角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param role body model.Role true "角色信息"
// @Success 200 {object} model.Role
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/roles [post]
func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

// GetRole 查询角色
// @Summary 查询角色
// @Description 根据ID查询角色
// @Tags 角色管理
// @Produce json
// @Param id query int true "角色ID"
// @Success 200 {object} model.Role
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/roles [get]
func GetRole(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数id无效"})
		return
	}
	var role model.Role
	if err := repository.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色未找到"})
		return
	}
	c.JSON(http.StatusOK, role)
}

// UpdateRole 更新角色
// @Summary 更新角色
// @Description 根据ID更新角色信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id query int true "角色ID"
// @Param role body model.Role true "角色信息"
// @Success 200 {object} model.Role
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/roles [put]
func UpdateRole(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数id无效"})
		return
	}
	var role model.Role
	if err := repository.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "角色未找到"})
		return
	}
	var req model.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role.RoleName = req.RoleName
	role.RoleCode = req.RoleCode
	if err := repository.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, role)
}

// DeleteRole 删除角色
// @Summary 删除角色
// @Description 根据ID删除角色
// @Tags 角色管理
// @Produce json
// @Param id query int true "角色ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/roles [delete]
func DeleteRole(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数id无效"})
		return
	}
	if err := repository.DB.Delete(&model.Role{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetUserBindRole 查询用户绑定角色id
// @Summary 查询用户绑定角色id
// @Description 根据用户id查询所有角色id
// @Tags 角色管理
// @Produce json
// @Param user_id query int true "用户ID"
// @Success 200 {object} []uint
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/roles/bind [get]
func GetUserBindRole(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数user_id无效"})
		return
	}

	roleIDs, err := repository.GetRoleIDsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询角色ID失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role_ids": roleIDs})
}
