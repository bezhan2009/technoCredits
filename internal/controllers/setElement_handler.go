package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/service"
)

// CreateSettlement godoc
// @Summary      Create settlement
// @Description  Create a new settlement record
// @Tags         settlements
// @Accept       json
// @Produce      json
// @Param        settlement  body      models.Settlement  true  "Settlement data"
// @Success      201  {object}  models.Settlement
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /settlements [post]
func CreateSettlement(c *gin.Context) {
	var settlement models.Settlement
	if err := c.ShouldBindJSON(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	if err := service.SettlementCreate(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, settlement)
}

// GetSettlementByID godoc
// @Summary      Get settlement by ID
// @Description  Get a settlement record by ID
// @Tags         settlements
// @Produce      json
// @Param        id   path      int  true  "Settlement ID"
// @Success      200  {object}  models.Settlement
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /settlements/{id} [get]
func GetSettlementByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, err)
		return
	}

	settlement, err := service.GetSettlementByID(uint(id))
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, settlement)
}

// GetAllSettlements godoc
// @Summary      Get all settlements
// @Description  Get a list of all settlements
// @Tags         settlements
// @Produce      json
// @Success      200  {array}   models.Settlement
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /settlements [get]
func GetAllSettlements(c *gin.Context) {
	settlements, err := service.GetAllSettlements()
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, settlements)
}

// UpdateSettlement godoc
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
// @Security ApiKeyAuth
// @Router       /settlements/{id} [put]
func UpdateSettlement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, err)
		return
	}

	var settlement models.Settlement
	if err = c.ShouldBindJSON(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	settlement.ID = uint(id)
	if err = service.UpdateSettlements(&settlement); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, settlement)
}

// DeleteSettlement godoc
// @Summary      Delete settlement
// @Description  Delete a settlement record by ID
// @Tags         settlements
// @Param        id   path      int  true  "Settlement ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /settlements/{id} [delete]
func DeleteSettlement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, err)
		return
	}

	if err := service.DeleteSettlement(uint(id)); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
