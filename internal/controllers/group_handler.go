package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/service"
	"technoCredits/internal/controllers/middlewares"
	"technoCredits/pkg/errs"
)

// CreateGroup godoc
// @Summary Создать группу
// @Description Создаёт новую группу
// @Tags groups
// @Accept json
// @Produce json
// @Param group body models.Group true "Группа"
// @Success 201 {object} models.Group
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups [post]
func CreateGroup(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		HandleError(c, err)
		return
	}

	group.OwnerID = userID

	if err := service.CreateGroup(group); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Group created successfully",
	})
}

// GetGroupsUser godoc
// @Summary Получить все группы пользователя
// @Description Возвращает все группы пользователя
// @Tags groups
// @Produce json
// @Success 200 {object} []models.Group
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups [get]
func GetGroupsUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	groups, err := service.GetAllUserGroups(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroupByID godoc
// @Summary Получить группу по ID
// @Description Возвращает группу по идентификатору
// @Tags groups
// @Produce json
// @Param group_id path int true "ID группы"
// @Success 200 {object} models.Group
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{group_id} [get]
func GetGroupByID(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		HandleError(c, err)
		return
	}

	group, err := service.GetAllUserGroupsByID(userID, uint(groupID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, group)
}

// UpdateGroup godoc
// @Summary Обновить группу
// @Description Обновляет данные группы
// @Tags groups
// @Accept json
// @Produce json
// @Param group_id path int true "ID группы"
// @Param group body models.Group true "Группа"
// @Success 200 {object} models.Group
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{group_id} [put]
func UpdateGroup(c *gin.Context) {
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var group models.Group
	if err = c.ShouldBindJSON(&group); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	group.ID = uint(groupID)

	if err = service.UpdateGroup(group); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, group)
}

// DeleteGroup godoc
// @Summary Удалить группу
// @Description Удаляет группу по ID
// @Tags groups
// @Param group_id path int true "ID группы"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{group_id} [delete]
func DeleteGroup(c *gin.Context) {
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		HandleError(c, err)
		return
	}

	if err = service.DeleteGroup(uint(groupID)); err != nil {
		HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
