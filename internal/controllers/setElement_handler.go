package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/service"
)

type SettlementHandler struct {
	service *service.SettlementService
}

func NewSettlementHandler(service *service.SettlementService) *SettlementHandler {
	return &SettlementHandler{service: service}
}

func (h *SettlementHandler) Create(c *gin.Context) {
	var settlement models.Settlement
	if err := c.ShouldBindJSON(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	if err := h.service.Create(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, settlement)
}

func (h *SettlementHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	settlement, err := h.service.GetByID(uint(id))
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, settlement)
}

func (h *SettlementHandler) GetAll(c *gin.Context) {
	settlements, err := h.service.GetAll()
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, settlements)
}

func (h *SettlementHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var settlement models.Settlement
	if err := c.ShouldBindJSON(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	settlement.ID = uint(id)
	if err := h.service.Update(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, settlement)
}

func (h *SettlementHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
