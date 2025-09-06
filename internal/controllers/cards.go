package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"technoCredits/internal/app/models"
	"technoCredits/internal/app/service"
	"technoCredits/internal/app/service/validators"
	"technoCredits/internal/controllers/middlewares"
	"technoCredits/pkg/errs"
)

// GetAllCardsUser godoc
// @Summary Получить все карты расходов пользователя
// @Description Возвращает список карт расходов пользователя за указанный месяц и год с возможностью поиска и пагинации
// @Tags Карты расходов
// @Accept  json
// @Produce  json
// @Param month query int true "Месяц (1-12)"
// @Param year query int true "Год (например: 2024)"
// @Param after query string false "ID для пагинации (получить записи после указанного ID)"
// @Param search query string false "Поиск по описанию или именам пользователей"
// @Security ApiKeyAuth
// @Success 200 {array} models.CardsExpense "Список карт расходов"
// @Router /cards [get]
func GetAllCardsUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	afterIDStr := c.Query("after")
	if afterIDStr == "" {
		afterIDStr = "0"
	}

	search := c.Query("after")

	afterID, err := strconv.Atoi(afterIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidAfterID)
		return
	}

	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidMonth)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidYear)
		return
	}

	cardsExpense, err := service.GetAllCardsUser(month, year, int(userID), afterID, search)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, cardsExpense)
}

// CreateCardExpense godoc
// @Summary Создать новую карту расходов
// @Description Создает новую карту расходов для группы пользователей
// @Tags Карты расходов
// @Accept  json
// @Produce  json
// @Param card body models.CardsExpense true "Данные карты расходов"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Карта расходов успешно создана"
// @Router /cards [post]
func CreateCardExpense(c *gin.Context) {
	var cardsExpense models.CardsExpense
	if err := c.ShouldBind(&cardsExpense); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateCardExpense(cardsExpense)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully created card expense",
	})
}

// CreateCardExpensePayer godoc
// @Summary Добавить плательщика к карте расходов
// @Description Добавляет пользователя как плательщика для определенной карты расходов
// @Tags Плательщики
// @Accept  json
// @Produce  json
// @Param payer body models.CardsExpensePayer true "Данные плательщика"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Плательщик успешно добавлен"
// @Router /cards/payers [post]
func CreateCardExpensePayer(c *gin.Context) {
	var cardsExpensePayer models.CardsExpensePayer
	if err := c.ShouldBind(&cardsExpensePayer); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateCardExpensePayer(cardsExpensePayer)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully created card expense payer",
	})
}

// CreateCardExpenseUser godoc
// @Summary Добавить участника к карте расходов
// @Description Добавляет пользователя как участника расходов для определенной карты
// @Tags Участники
// @Accept  json
// @Produce  json
// @Param user body models.CardsExpenseUser true "Данные участника"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Участник успешно добавлен"
// @Router /cards/users [post]
func CreateCardExpenseUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	var cardsExpenseUser models.CardsExpenseUser
	if err := c.ShouldBind(&cardsExpenseUser); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateCardExpenseUser(cardsExpenseUser, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully created card expense user",
	})
}

// UpdateCardExpense godoc
// @Summary Обновить карту расходов
// @Description Обновляет информацию о карте расходов
// @Tags Карты расходов
// @Accept  json
// @Produce  json
// @Param id path int true "ID карты расходов"
// @Param card body models.CardsExpense true "Обновленные данные карты расходов"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Карта расходов успешно обновлена"
// @Router /cards/{id} [patch]
func UpdateCardExpense(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var cardsExpense models.CardsExpense
	if err = c.ShouldBind(&cardsExpense); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	cardsExpense.ID = uint(cardID)
	err = service.UpdateCardExpense(cardsExpense)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated card expense",
	})
}

// UpdateCardExpensePayer godoc
// @Summary Обновить информацию о плательщике
// @Description Обновляет данные плательщика для карты расходов
// @Tags Плательщики
// @Accept  json
// @Produce  json
// @Param id path int true "ID карты расходов плательщика"
// @Param payer body models.CardsExpensePayer true "Обновленные данные плательщика"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Плательщик успешно обновлен"
// @Router /cards/payers/{id} [patch]
func UpdateCardExpensePayer(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var cardsExpensePayer models.CardsExpensePayer
	if err := c.ShouldBind(&cardsExpensePayer); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	cardsExpensePayer.ID = uint(cardID)

	err = service.UpdateCardExpensePayer(cardsExpensePayer)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated card expense payer",
	})
}

// UpdateCardExpenseUser godoc
// @Summary Обновить информацию об участнике
// @Description Обновляет данные участника карты расходов
// @Tags Участники
// @Accept  json
// @Produce  json
// @Param id path int true "ID карты расходов участнике"
// @Param user body models.CardsExpenseUser true "Обновленные данные участника"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Участник успешно обновлен"
// @Router /cards/users/{id} [patch]
func UpdateCardExpenseUser(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var cardsExpenseUser models.CardsExpenseUser
	if err := c.ShouldBind(&cardsExpenseUser); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	cardsExpenseUser.ID = uint(cardID)

	err = service.UpdateCardExpenseUser(cardsExpenseUser)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated card expense user",
	})
}

// DeleteCardExpense godoc
// @Summary Удалить карту расходов
// @Description Удаляет карту расходов по ID
// @Tags Карты расходов
// @Accept  json
// @Produce  json
// @Param id path int true "ID карты расходов"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Карта расходов успешно удалена"
// @Router /cards/{id} [delete]
func DeleteCardExpense(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteCardExpense(uint(cardID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted card expense",
	})
}

// DeleteCardExpensePayer godoc
// @Summary Удалить плательщика
// @Description Удаляет плательщика по ID
// @Tags Плательщики
// @Accept  json
// @Produce  json
// @Param id path int true "ID плательщика"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Плательщик успешно удален"
// @Router /cards/payers/{id} [delete]
func DeleteCardExpensePayer(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteCardExpensePayer(uint(cardID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted card expense payer",
	})
}

// DeleteCardExpenseUser godoc
// @Summary Удалить участника
// @Description Удаляет участника карты расходов по ID
// @Tags Участники
// @Accept  json
// @Produce  json
// @Param id path int true "ID участника"
// @Security ApiKeyAuth
// @Success 200 {object} models.DefaultResponse "Участник успешно удален"
// @Router /cards/users/{id} [delete]
func DeleteCardExpenseUser(c *gin.Context) {
	cardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteCardExpenseUser(uint(cardID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete card expense user",
	})
}
