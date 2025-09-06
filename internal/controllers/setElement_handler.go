package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/service"
)

// SettlementHandler обрабатывает CRUD для Settlement
type SettlementHandler struct {
	service *service.SettlementService
}

func NewSettlementHandler(service *service.SettlementService) *SettlementHandler {
	return &SettlementHandler{service: service}
}

// Create Settlement
// @Summary      Create settlement
// @Description  Create a new settlement record
// @Tags         settlements
// @Accept       json
// @Produce      json
// @Param        settlement  body      models.Settlement  true  "Settlement data"
// @Success      201  {object}  models.Settlement
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /settlements [post]
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

// GetByID Settlement
// @Summary      Get settlement by ID
// @Description  Get a settlement record by ID
// @Tags         settlements
// @Produce      json
// @Param        id   path      int  true  "Settlement ID"
// @Success      200  {object}  models.Settlement
// @Failure      404  {object}  map[string]string
// @Router       /settlements/{id} [get]
func (h *SettlementHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	settlement, err := h.service.GetByID(uint(id))
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, settlement)
}

// GetAll Settlements
// @Summary      Get all settlements
// @Description  Get a list of all settlements
// @Tags         settlements
// @Produce      json
// @Success      200  {array}   models.Settlement
// @Failure      500  {object}  map[string]string
// @Router       /settlements [get]
func (h *SettlementHandler) GetAll(c *gin.Context) {
	settlements, err := h.service.GetAll()
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, settlements)
}

// Update Settlement
// @Summary      Update settlement
// @Description  Update a settlement record
// @Tags         settlements
// @Accept       json
// @Produce      json
// @Param        id         path      int              true  "Settlement ID"
// @Param        settlement body      models.Settlement true "Settlement data"
// @Success      200  {object}  models.Settlement
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /settlements/{id} [put]
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

// Delete Settlement
// @Summary      Delete settlement
// @Description  Delete a settlement record by ID
// @Tags         settlements
// @Param        id   path      int  true  "Settlement ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /settlements/{id} [delete]
func (h *SettlementHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
