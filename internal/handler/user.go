package handler

import (
	"fuseops-backend/internal/model"
	"fuseops-backend/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser 创建用户
// @Summary 创建用户
// @Description 新增一个用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body model.User true "用户信息"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/users [post]
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 密码加密
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}
	user.Password = string(hashedPwd)
	if err := repository.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// GetUser 查询用户
// @Summary 查询用户
// @Description 根据ID查询用户
// @Tags 用户管理
// @Produce json
// @Param id query int true "用户ID"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/users [get]
func GetUser(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数id无效"})
		return
	}
	var user model.User
	if err := repository.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Description 根据ID更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id query int true "用户ID"
// @Param user body model.User true "用户信息"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/users [put]
func UpdateUser(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数id无效"})
		return
	}
	var user model.User
	if err := repository.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 如果密码有修改，进行加密
	if req.Password != "" && req.Password != user.Password {
		hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
			return
		}
		user.Password = string(hashedPwd)
	}
	// 其他字段更新
	user.UserName = req.UserName
	user.RealName = req.RealName
	user.Avatar = req.Avatar
	user.RoleID = req.RoleID
	user.GradeID = req.GradeID
	user.Status = req.Status

	if err := repository.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 根据ID删除用户
// @Tags 用户管理
// @Produce json
// @Param id query int true "用户ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/users [delete]
func DeleteUser(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数id无效"})
		return
	}
	if err := repository.DB.Delete(&model.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
