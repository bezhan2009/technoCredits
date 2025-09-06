package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technoCredits/internal/app/service"
)

type GroupMemberHandler struct {
	Service *service.GroupMemberService
}

func NewGroupMemberHandler(s *service.GroupMemberService) *GroupMemberHandler {
	return &GroupMemberHandler{Service: s}
}

// AddMember godoc
// @Summary Добавить участника в группу
// @Description Добавляет нового участника в группу с ролью
// @Tags group-members
// @Accept json
// @Produce json
// @Param groupID path int true "ID группы"
// @Param member body object{user_id=int,role=string} true "Данные участника"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{groupID}/members [post]
func (h *GroupMemberHandler) AddMember(c *gin.Context) {
	groupID, _ := strconv.Atoi(c.Param("groupID"))

	var req struct {
		UserID uint   `json:"user_id"`
		Role   string `json:"role"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.Service.AddMember(uint(groupID), req.UserID, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "member added"})
}

// GetMembers godoc
// @Summary Получить участников группы
// @Description Возвращает список всех участников группы
// @Tags group-members
// @Produce json
// @Param groupID path int true "ID группы"
// @Success 200 {array} object{group_id=int,user_id=int,role=string,joined_at=string}
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{groupID}/members [get]
func (h *GroupMemberHandler) GetMembers(c *gin.Context) {
	groupID, _ := strconv.Atoi(c.Param("groupID"))

	members, err := h.Service.GetMembers(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

// UpdateMemberRole godoc
// @Summary Обновить роль участника
// @Description Обновляет роль конкретного участника группы
// @Tags group-members
// @Accept json
// @Produce json
// @Param groupID path int true "ID группы"
// @Param userID path int true "ID пользователя"
// @Param role body object{role=string} true "Новая роль"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{groupID}/members/{userID} [put]
func (h *GroupMemberHandler) UpdateMemberRole(c *gin.Context) {
	groupID, _ := strconv.Atoi(c.Param("groupID"))
	userID, _ := strconv.Atoi(c.Param("userID"))

	var req struct {
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.Service.UpdateMemberRole(uint(groupID), uint(userID), req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role updated"})
}

// RemoveMember godoc
// @Summary Удалить участника из группы
// @Description Удаляет участника по ID из группы
// @Tags group-members
// @Produce json
// @Param groupID path int true "ID группы"
// @Param userID path int true "ID пользователя"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /groups/{groupID}/members/{userID} [delete]
func (h *GroupMemberHandler) RemoveMember(c *gin.Context) {
	groupID, _ := strconv.Atoi(c.Param("groupID"))
	userID, _ := strconv.Atoi(c.Param("userID"))

	err := h.Service.RemoveMember(uint(groupID), uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member removed"})
}
