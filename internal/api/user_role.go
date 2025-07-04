package api

import (
	"fuseops-backend/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
