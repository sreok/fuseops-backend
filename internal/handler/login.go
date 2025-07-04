package handler

import (
	"fuseops-backend/internal/model"
	"fuseops-backend/internal/pkg"
	"fuseops-backend/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest 登录请求结构体
type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户名和密码登录
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param login body LoginRequest true "登录信息"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "参数错误"})
		return
	}
	var user model.User
	// 查询用户名
	if err := repository.DB.Where("username = ?", req.UserName).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "用户名或密码错误"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "用户名或密码错误"})
		return
	}

	// 查询角色所有ID
	roleIDs, err := repository.GetRoleIDsByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询角色ID失败"})
		return
	}

	// 根据角色ID查询所有角色代码
	roleCodes, err := repository.GetRoleCodeByRoleIDs(roleIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询角色ID失败"})
		return
	}

	// 生成token
	token, expires, err := pkg.GenerateToken(user.ID, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"avatar":       user.Avatar,
			"username":     user.UserName,
			"nickname":     user.RealName,
			"roles":        roleCodes,
			"permissions":  []string{"permission:btn:add", "permission:btn:edit"},
			"accessToken":  token,
			"refreshToken": token + "Refresh",
			"expires":      expires,
		},
	})
}
