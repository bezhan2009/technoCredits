package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/service"
	"technoCredits/pkg/errs"
)

// GetGroupMembersByGroupID godoc
// @Summary Получить участников группы
// @Description Возвращает всех участников группы по ID группы
// @Tags group-members
// @Produce json
// @Param group_id path int true "ID группы"
// @Success 200 {object} models.GroupMember
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{group_id}/members [get]
func GetGroupMembersByGroupID(c *gin.Context) {
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	members, err := service.GetGroupMembersByGroupID(uint(groupID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, members)
}

// CreateGroupMember godoc
// @Summary Добавить участника в группу
// @Description Добавляет нового участника в группу
// @Tags group-members
// @Accept json
// @Produce json
// @Param group_id path int true "ID группы"
// @Param member body models.GroupMember true "Участник группы"
// @Success 201 {object} models.GroupMember
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{group_id}/members [post]
func CreateGroupMember(c *gin.Context) {
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	var member models.GroupMember
	if err = c.ShouldBindJSON(&member); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	member.GroupID = uint(groupID)

	if err = service.CreateGroupMember(member); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Group member created",
	})
}

// UpdateGroupMember godoc
// @Summary Обновить участника группы
// @Description Обновляет данные участника группы
// @Tags group-members
// @Accept json
// @Produce json
// @Param id path int true "ID участника группы"
// @Param member body models.GroupMember true "Участник группы"
// @Success 200 {object} models.GroupMember
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /group-members/{id} [put]
func UpdateGroupMember(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var member models.GroupMember
	if err = c.ShouldBindJSON(&member); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	member.ID = uint(id)

	if err = service.UpdateGroupMember(member); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, member)
}

// DeleteGroupMember godoc
// @Summary Удалить участника из группы
// @Description Удаляет участника из группы по ID
// @Tags group-members
// @Param id path int true "ID участника группы"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /group-members/{id} [delete]
func DeleteGroupMember(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		HandleError(c, err)
		return
	}

	if err = service.DeleteGroupMember(uint(id)); err != nil {
		HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
